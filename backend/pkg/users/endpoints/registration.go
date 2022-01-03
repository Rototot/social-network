package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"social-network/pkg/users"
	"social-network/pkg/users/services"
)

type PostRegistrationRequest struct {
	services.RegisterParams
}

type PostRegistrationResponse struct {
	Id users.UserID
}

func MakePostRegistrationsEndpoint(service services.RegisterServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostRegistrationRequest)

		user, err := service.Register(ctx, services.RegisterParams{
			Email:     req.Email,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Age:       req.Age,
			Gender:    req.Gender,
			City:      req.City,
			Interests: req.Interests,
		})

		if err != nil {
			return nil, err
		}

		return &PostRegistrationResponse{
			Id: user.ID,
		}, nil
	}
}
