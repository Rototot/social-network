package services

import (
	"context"
)

type LogoutServiceInterface interface {
	Logout(ctx context.Context, id SessionId) error
}

type LogoutService struct {
	sessions SessionStorageInterface
}

func NewLogoutService(sessions SessionStorageInterface) *LogoutService {
	return &LogoutService{sessions: sessions}
}

func (s *LogoutService) Logout(ctx context.Context, id SessionId) error {
	return s.sessions.Remove(ctx, id)
}
