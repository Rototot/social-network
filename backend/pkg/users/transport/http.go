package transport

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	commonHttp "social-network/pkg/common/transport/http"
	usersEndpoints "social-network/pkg/users/endpoints"
)

func NewHttpHandler(
	router *mux.Router,
	endpoints usersEndpoints.Endpoints,
	logger log.Logger,
) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(commonHttp.MakeErrorEncoder()),
	}

	router.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.PostLogin,
		commonHttp.MakeRequestDecoder(func() interface{} {
			return &usersEndpoints.PostLoginRequest{}
		}),
		commonHttp.MakeResponseEncoder(),
		options...,
	))

	router.Methods("POST").Path("/register").Handler(httptransport.NewServer(
		endpoints.PostRegistration,
		commonHttp.MakeRequestDecoder(func() interface{} {
			return &usersEndpoints.PostRegistrationRequest{}
		}),
		commonHttp.MakeResponseEncoder(),
		options...,
	))

	router.Methods("POST").Path("/logout").Handler(httptransport.NewServer(
		endpoints.PostLogin,
		commonHttp.MakeRequestDecoder(func() interface{} {
			return &usersEndpoints.PostLogoutRequest{}
		}),
		commonHttp.MakeResponseEncoder(),
		options...,
	))

	return router
}
