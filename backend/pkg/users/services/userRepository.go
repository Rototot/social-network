package services

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./userRepository_mocked.go

import (
	"context"
	"social-network/pkg/users"
)

type UserRepositoryInterface interface {
	FindByEmailAndPassword(cxt context.Context, email string, password users.HashedPassword) (*users.User, error)
}
