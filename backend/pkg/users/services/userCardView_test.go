package services

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"social-network/pkg/users"
	"testing"
)

func factoryViewService(t *testing.T) (
	*UserCardViewer,
	*gomock.Controller,
	*MockUserRepositoryInterface,
) {
	ctrl := gomock.NewController(t)
	uRepository := NewMockUserRepositoryInterface(ctrl)

	service := NewUserCardViewer(
		uRepository,
	)

	return service, ctrl, uRepository
}

func TestViewWhenUserExistsThenOk(t *testing.T) {
	service, controller, repository := factoryViewService(t)
	defer controller.Finish()

	cxt := context.TODO()
	expectedId := users.UserID(10)
	expectedUser := &users.User{
		ID: expectedId,
	}

	repository.EXPECT().
		FindById(cxt, expectedId).
		Return(expectedUser, nil).
		Times(1)

	result, err := service.View(cxt, expectedId)

	assert.Nil(t, err)
	assert.Equal(t, result, expectedUser)
}

func TestViewWhenUserNotExistsThenError(t *testing.T) {
	service, controller, repository := factoryViewService(t)
	defer controller.Finish()

	cxt := context.TODO()
	expectedId := users.UserID(10)

	repository.EXPECT().
		FindById(cxt, expectedId).
		Return(nil, nil).
		Times(1)

	result, err := service.View(cxt, expectedId)

	assert.Equal(t, err, users.ErrUserNotFound)
	assert.Nil(t, result)
}

func TestViewWhenErrorThenError(t *testing.T) {
	service, controller, repository := factoryViewService(t)
	defer controller.Finish()

	cxt := context.TODO()
	expectedId := users.UserID(10)
	expectedError := errors.New("test error")

	repository.EXPECT().
		FindById(cxt, expectedId).
		Return(nil, expectedError).
		Times(1)

	result, err := service.View(cxt, expectedId)

	assert.Equal(t, err, expectedError)
	assert.Nil(t, result)
}
