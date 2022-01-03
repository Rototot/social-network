package transport

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	commonHttp "social-network/pkg/common/transport/http"
	"social-network/pkg/users"
	"social-network/pkg/users/endpoints"
)

func NewHttpHandler(
	endpoints endpoints.Endpoints,
	logger log.Logger,
) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(commonHttp.MakeErrorEncoder()),
	}

	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.PostLogin,
		commonHttp.MakeRequestDecoder(func() any {
			return users.LoginParams{}
		}),
		commonHttp.MakeResponseEncoder(),
		options...,
	))

	r.Methods("POST").Path("/register").Handler(httptransport.NewServer(
		endpoints.PostRegistration,
		commonHttp.MakeRequestDecoder(func() any {
			return users.RegisterParams{}
		}),
		commonHttp.MakeResponseEncoder(),
		options...,
	))

	r.Methods("POST").Path("/request-reset-password").Handler(httptransport.NewServer(
		endpoints.PostRequestResetPassword,
		commonHttp.MakeRequestDecoder(func() any {
			return users.RequestResetPasswordParams{}
		}),
		commonHttp.MakeResponseEncoder(),
		options...,
	))

	r.Methods("POST").Path("/reset-password").Handler(httptransport.NewServer(
		endpoints.PostResetPassword,
		commonHttp.MakeRequestDecoder(func() any {
			return users.ResetPasswordParams{}
		}),
		commonHttp.MakeResponseEncoder(),
		options...,
	))

	r.Methods("POST").Path("/logout").Handler(httptransport.NewServer(
		endpoints.PostLogin,
		commonHttp.MakeRequestDecoder(func() any {
			return users.ResetPasswordParams{}
		}),
		commonHttp.MakeResponseEncoder(),
		options...,
	))

	return r
}
