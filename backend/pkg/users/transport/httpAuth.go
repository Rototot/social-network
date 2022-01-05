package transport

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"net/http"
	commonHttp "social-network/pkg/common/transport/http"
	usersEndpoints "social-network/pkg/users/endpoints"
	"social-network/pkg/users/middleware"
)

func MakeAuthHttpHandler(
	endpoints usersEndpoints.Endpoints,
	logger log.Logger,
) http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/auth").Subrouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(commonHttp.MakeErrorEncoder(extendErrorStatusMap)),
	}

	s.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.PostLogin,
		commonHttp.MakeRequestDecoder(func(_ context.Context, r *http.Request) (interface{}, error) {
			req := &usersEndpoints.PostLoginRequest{}
			err := easyjson.UnmarshalFromReader(r.Body, req)

			return req, err
		}),
		commonHttp.MakeResponseEncoder(extendErrorStatusMap),
		options...,
	))

	s.Methods("POST").Path("/register").Handler(httptransport.NewServer(
		endpoints.PostRegistration,
		commonHttp.MakeRequestDecoder(func(_ context.Context, r *http.Request) (interface{}, error) {
			req := &usersEndpoints.PostRegistrationRequest{}
			err := easyjson.UnmarshalFromReader(r.Body, req)

			return req, err
		}),
		commonHttp.MakeResponseEncoder(extendErrorStatusMap),
		options...,
	))

	s.Methods("POST").Path("/logout").Handler(httptransport.NewServer(
		endpoints.PostLogin,
		commonHttp.MakeRequestDecoder(func(ctx context.Context, r *http.Request) (interface{}, error) {
			return &usersEndpoints.PostLogoutRequest{
				SessionId: middleware.ExtractToken(ctx),
			}, nil
		}),
		commonHttp.MakeResponseEncoder(extendErrorStatusMap),
		options...,
	))

	return r
}
