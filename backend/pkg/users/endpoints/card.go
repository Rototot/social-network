package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"social-network/pkg/users"
	"social-network/pkg/users/services"
)

type GetUserCardRequest struct {
	Id users.UserID
}

type GetUserCardResponse struct {
	ID        users.UserID
	Email     string
	Password  users.HashedPassword
	FirstName string
	LastName  string
	Age       int8
	Gender    users.Gender
	City      string
	Interests []string
}

func MakeGetUserCardEndpoint(repository services.UserRepositoryInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserCardRequest)

		user, err := repository.FindById(ctx, req.Id)
		if err != nil {
			return nil, err
		}

		return &GetUserCardResponse{
			ID:        user.ID,
			Email:     user.Email,
			Password:  user.Password,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Age:       user.Age,
			Gender:    user.Gender,
			City:      user.City,
			Interests: user.Interests,
		}, nil
	}
}
