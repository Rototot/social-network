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
	endpoints usersEndpoints.Endpoints,
	logger log.Logger,
) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(commonHttp.MakeErrorEncoder()),
	}

	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.PostLogin,
		commonHttp.MakeRequestDecoder(func() interface{} {
			return &usersEndpoints.PostLoginRequest{}
		}),
		commonHttp.MakeResponseEncoder(),
		options...,
	))

	r.Methods("POST").Path("/register").Handler(httptransport.NewServer(
		endpoints.PostRegistration,
		commonHttp.MakeRequestDecoder(func() interface{} {
			return &usersEndpoints.PostRegistrationRequest{}
		}),
		commonHttp.MakeResponseEncoder(),
		options...,
	))

	r.Methods("POST").Path("/logout").Handler(httptransport.NewServer(
		endpoints.PostLogin,
		commonHttp.MakeRequestDecoder(func() interface{} {
			return &usersEndpoints.PostLogoutRequest{}
		}),
		commonHttp.MakeResponseEncoder(),
		options...,
	))

	return r
}
