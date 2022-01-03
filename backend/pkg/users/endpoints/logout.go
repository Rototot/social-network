package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"social-network/pkg/users/services"
)

type PostLogoutRequest struct {
	SessionId services.SessionId
}

func MakePostLogoutEndpoint(service services.LogoutServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostLogoutRequest)

		err := service.Logout(ctx, req.SessionId)

		return nil, err
	}
}
