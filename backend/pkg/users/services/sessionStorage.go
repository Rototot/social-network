package services

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./sessionStorage_mocked.go

import (
	"context"
	"social-network/pkg/users"
	"time"
)

type SessionId string

type SessionStorageInterface interface {
	Add(ctx context.Context, id users.UserID, expire time.Duration) (SessionId, error)
	//Has(id users.UserID) (bool, error)
	Remove(ctx context.Context, id SessionId) error
	Get(ctx context.Context, id SessionId) (users.UserID, error)
}
