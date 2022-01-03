package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"social-network/pkg/users"
	"testing"
	"time"
)

func factoryLoginService(t *testing.T) (
	*LoginService,
	*gomock.Controller,
	*MockUserRepositoryInterface,
	*MockPasswordHasher,
	*MockSessionStorageInterface,
	time.Duration,
) {
	ctrl := gomock.NewController(t)

	uRepository := NewMockUserRepositoryInterface(ctrl)
	hasher := NewMockPasswordHasher(ctrl)
	sessions := NewMockSessionStorageInterface(ctrl)
	expire := time.Minute

	service := NewLoginService(
		uRepository,
		hasher,
		sessions,
		expire,
	)

	return service, ctrl, uRepository, hasher, sessions, expire
}

func TestLoginWhenIsOkThenSetSession(t *testing.T) {
	service, controller, userRepository, passwordHasher, sessionStorage, duration := factoryLoginService(t)
	defer controller.Finish()

	ctx := context.TODO()

	payload := LoginParams{
		Email:    "test@email.local",
		Password: "test password",
	}
	expectedPasswordHash := users.HashedPassword("test_password_hash_123")
	expectedUser := &users.User{
		ID:        10,
		Email:     payload.Email,
		Password:  expectedPasswordHash,
		FirstName: "First Name",
		LastName:  "Last Name",
		Age:       20,
		Gender:    users.Female,
		City:      "Moscow",
		Interests: nil,
	}
	expectedSessionId := SessionId("test_session_id")

	userRepository.EXPECT().
		FindByEmailAndPassword(ctx, payload.Email, expectedPasswordHash).
		Return(expectedUser, nil)

	passwordHasher.EXPECT().
		Hash(payload.Password).
		Return(expectedPasswordHash, nil)

	sessionStorage.EXPECT().
		Add(expectedUser.ID, duration).
		Return(expectedSessionId, nil)

	result, err := service.Login(ctx, payload)

	assert.Equal(t, result, expectedSessionId)
	assert.Nil(t, err)
}

func TestLoginWhenNotFoundUserThenError(t *testing.T) {
	service, controller, userRepository, passwordHasher, _, _ := factoryLoginService(t)
	defer controller.Finish()

	ctx := context.TODO()

	payload := LoginParams{
		Email:    "test@email.local",
		Password: "test password",
	}
	expectedPasswordHash := users.HashedPassword("test_password_hash_123")

	userRepository.EXPECT().
		FindByEmailAndPassword(ctx, payload.Email, expectedPasswordHash).
		Return(nil, users.ErrUserNotFound)

	passwordHasher.EXPECT().
		Hash(payload.Password).
		Return(expectedPasswordHash, nil)

	result, err := service.Login(ctx, payload)

	assert.Equal(t, result, "")
	assert.ErrorIs(t, err, users.ErrUserNotFound)
}
