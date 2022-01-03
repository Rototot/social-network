package services

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./userRepository_mocked.go

import (
	"context"
	"social-network/pkg/users"
)

type UserRepositoryInterface interface {
	FindByEmailAndPassword(ctx context.Context, email string, password users.HashedPassword) (*users.User, error)

	FindById(ctx context.Context, id users.UserID) (*users.User, error)
}
