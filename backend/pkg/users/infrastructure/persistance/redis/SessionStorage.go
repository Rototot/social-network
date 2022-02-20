package redis

import (
	"context"
	"crypto/sha512"
	"fmt"
	"social-network/pkg/users"
	"social-network/pkg/users/services"
	"strconv"
	"strings"
	"time"
)
import redis "github.com/go-redis/redis/v8"

const sessionKeyPrefix = "sessions:users:"
const sessionKeyPartsSeparator = ":"

type SessionRepository struct {
	client *redis.Client
}

func NewSessionRepository(client *redis.Client) *SessionRepository {
	return &SessionRepository{client: client}
}

func (r *SessionRepository) Add(ctx context.Context, id users.UserID, expire time.Duration) (services.SessionId, error) {
	// todo добавить сложности генерации сессии
	session, err := r.generateSession(id)
	if err != nil {
		return "", err
	}

	key := r.generateKeyName(string(session))
	err = r.client.SetXX(ctx, key, id, expire).Err()
	if err != nil {
		return "", err
	}

	return session, nil
}

func (r *SessionRepository) Remove(ctx context.Context, id services.SessionId) error {
	key := r.generateKeyName(string(id))

	return r.client.Del(ctx, key).Err()
}

func (r *SessionRepository) Get(ctx context.Context, id services.SessionId) (users.UserID, error) {
	key := r.generateKeyName(string(id))

	rawUserID, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	userId, err := strconv.Atoi(rawUserID)
	if err != nil {
		return 0, err
	}

	return users.UserID(userId), nil
}

func (r *SessionRepository) generateKeyName(value ...string) string {
	parts := append([]string{sessionKeyPrefix}, value...)

	return strings.Join(parts, sessionKeyPartsSeparator)
}

func (r *SessionRepository) generateSession(value users.UserID) (services.SessionId, error) {
	hasher := sha512.New()
	_, err := hasher.Write([]byte(value.String()))
	if err != nil {
		return "", err
	}

	return services.SessionId(fmt.Sprintf("%x", hasher.Sum(nil))), nil
}
