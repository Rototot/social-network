package services

import (
	"context"
	"social-network/pkg/users"
	"time"
)

type LoginParams struct {
	Email    string
	Password string
}

type LoginServiceInterface interface {
	Login(ctx context.Context, payload LoginParams) (SessionId, error)
}

type LoginService struct {
	repository    UserRepositoryInterface
	hasher        PasswordHasher
	sessions      SessionStorageInterface
	sessionExpire time.Duration
}

func NewLoginService(
	repository UserRepositoryInterface,
	hasher PasswordHasher,
	sessions SessionStorageInterface,
	sessionExpire time.Duration,
) *LoginService {
	return &LoginService{
		repository:    repository,
		hasher:        hasher,
		sessions:      sessions,
		sessionExpire: sessionExpire,
	}
}

func (h *LoginService) Login(ctx context.Context, payload LoginParams) (SessionId, error) {
	password, err := h.hasher.Hash(payload.Password)
	if err != nil {
		return "", err
	}

	user, err := h.repository.FindByEmailAndPassword(ctx, payload.Email, password)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", users.ErrUserNotFound
	}

	session, err := h.sessions.Add(user.ID, h.sessionExpire)
	if err != nil {
		return "", err
	}

	return session, nil
}
