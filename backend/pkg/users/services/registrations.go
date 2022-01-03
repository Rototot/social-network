package services

import (
	"context"
	"social-network/pkg/users"
)

type RegisterParams struct {
	Email     string
	FirstName string
	LastName  string
	Age       int8
	Gender    users.Gender
	City      string
	Interests []string
}

type RegisterServiceInterface interface {
	Register(ctx context.Context, payload RegisterParams) (*users.User, error)
}
