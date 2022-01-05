package endpoints

//go:generate easyjson -all $GOFILE

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"social-network/pkg/users/services"
)

type PostLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PostLoginResponse struct {
	Token services.SessionId `json:"token"`
}

func MakePostLoginEndpoint(service services.LoginServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*PostLoginRequest)

		token, err := service.Login(ctx, services.LoginParams{
			Email:    req.Email,
			Password: req.Password,
		})

		if err != nil {
			return nil, err
		}

		return &PostLoginResponse{
			Token: token,
		}, nil
	}
}
