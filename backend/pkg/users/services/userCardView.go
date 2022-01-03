package services

import (
	"context"
	"social-network/pkg/users"
)

type UserCardViewerInterface interface {
	View(id users.UserID) (*users.User, error)
}

type UserCardViewer struct {
	repository UserRepositoryInterface
}

func NewUserCardViewer(repository UserRepositoryInterface) *UserCardViewer {
	return &UserCardViewer{repository: repository}
}

func (s *UserCardViewer) View(ctx context.Context, id users.UserID) (*users.User, error) {
	user, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, users.ErrUserNotFound
	}

	return user, nil
}
