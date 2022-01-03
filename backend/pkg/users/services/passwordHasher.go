package services

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./passwordHasher_mocked.go

import (
	"crypto/sha512"
	"fmt"
	"social-network/pkg/users"
)

type PasswordHasherInterface interface {
	Hash(value string) (users.HashedPassword, error)
}

type PasswordHasher struct {
}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

func (p PasswordHasher) Hash(value string) (users.HashedPassword, error) {
	hasher := sha512.New()
	_, err := hasher.Write([]byte(value))
	if err != nil {
		return "", err
	}

	return users.HashedPassword(fmt.Sprintf("sha512:\t\t%x\n", hasher.Sum(nil))), nil
}
