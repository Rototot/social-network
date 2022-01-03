package services

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./passwordHasher_mocked.go

import (
	"social-network/pkg/users"
)

type PasswordHasher interface {
	Hash(value string) (users.HashedPassword, error)
}
