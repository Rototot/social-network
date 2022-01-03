package services

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./sessionStorage_mocked.go

import (
	"social-network/pkg/users"
	"time"
)

type SessionId string

type SessionStorageInterface interface {
	Add(id users.UserID, expire time.Duration) (SessionId, error)
	//Has(id users.UserID) (bool, error)
	Remove(SessionId) error
	//Get(id SessionId) (users.UserID, error)
}
