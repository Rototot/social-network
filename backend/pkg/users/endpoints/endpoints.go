package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"social-network/pkg/users/services"
)

type Endpoints struct {
	PostLogin        endpoint.Endpoint
	PostRegistration endpoint.Endpoint
	PostLogout       endpoint.Endpoint
}

func MakeEndpoints(
	loginService services.LoginServiceInterface,
	registrationService services.RegisterServiceInterface,
	logoutService services.LogoutServiceInterface,
) (Endpoints, error) {
	return Endpoints{
		PostLogin:        MakePostLoginEndpoint(loginService),
		PostRegistration: MakePostRegistrationsEndpoint(registrationService),
		PostLogout:       MakePostLogoutEndpoint(logoutService),
	}, nil
}
