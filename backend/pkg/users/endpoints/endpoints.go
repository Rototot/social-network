package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"social-network/pkg/users/middleware"
	"social-network/pkg/users/services"
)

type Endpoints struct {
	GetUserCard      endpoint.Endpoint
	PostLogin        endpoint.Endpoint
	PostRegistration endpoint.Endpoint
	PostLogout       endpoint.Endpoint
}

func MakeEndpoints(
	loginService services.LoginServiceInterface,
	registrationService services.RegisterServiceInterface,
	logoutService services.LogoutServiceInterface,
	userRepository services.UserRepositoryInterface,
	sessions services.SessionStorageInterface,

) (Endpoints, error) {
	authUser := middleware.AuthUserByToken(sessions)
	authGuard := middleware.AuthGuardMiddleware()
	guestGuards := middleware.OnlyGuestGuardMiddleware()

	return Endpoints{
		PostLogin:        guestGuards(authUser(MakePostLoginEndpoint(loginService))),
		PostRegistration: guestGuards(authUser(MakePostRegistrationsEndpoint(registrationService))),
		PostLogout:       authGuard(authUser(MakePostLogoutEndpoint(logoutService))),
		GetUserCard:      authGuard(authUser(MakeGetUserCardEndpoint(userRepository))),
	}, nil
}
