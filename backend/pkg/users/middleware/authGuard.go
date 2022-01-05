package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"social-network/pkg/users"
	"social-network/pkg/users/services"
)

const KeyUserIdKey = "user"
const KeyAuthTokenKey = "authToken"

func ExtractUserId(ctx context.Context) users.UserID {
	userId, ok := ctx.Value(KeyUserIdKey).(users.UserID)
	if !ok {
		return 0
	}

	return userId
}

func ExtractToken(ctx context.Context) services.SessionId {
	userId, ok := ctx.Value(KeyAuthTokenKey).(services.SessionId)
	if !ok {
		return ""
	}

	return userId
}

// AuthGuardMiddleware guard for access only auth users
func AuthGuardMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			userId, ok := ctx.Value(KeyUserIdKey).(users.UserID)
			if !ok || userId <= 0 {
				return 0, users.ErrAuthRequired
			}

			return next(ctx, request)
		}
	}
}

// AuthGuardMiddleware guard for access only auth users
func OnlyGuestGuardMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			userId, ok := ctx.Value(KeyUserIdKey).(users.UserID)
			if ok && userId > 0 {
				return nil, users.ErrPermissionDenied
			}

			return next(ctx, request)
		}
	}
}

// AuthUserByToken find user by session token
func AuthUserByToken(
	sessions services.SessionStorageInterface,
) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			token, ok := ctx.Value(KeyAuthTokenKey).(services.SessionId)
			if !ok || token == "" {
				ctx = context.WithValue(ctx, KeyUserIdKey, -1)
				return next(ctx, request)
			}

			// find session
			userId, err := sessions.Get(ctx, token)
			if err != nil || userId == 0 {
				return nil, users.ErrPermissionDenied
			}

			existsUserId, ok := ctx.Value(KeyUserIdKey).(users.UserID)
			if !ok || existsUserId == 0 {
				ctx = context.WithValue(ctx, KeyUserIdKey, userId)
			}

			return next(ctx, request)
		}
	}
}
