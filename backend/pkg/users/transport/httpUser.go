package transport

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	commonHttp "social-network/pkg/common/transport/http"
	usersEndpoints "social-network/pkg/users/endpoints"
	"social-network/pkg/users/middleware"
)

func MakeUserHttpHandler(
	endpoints usersEndpoints.Endpoints,
	logger log.Logger,
) http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/user").Subrouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(commonHttp.MakeErrorEncoder(extendErrorStatusMap)),
	}

	s.Methods("GET").Path("/card").Handler(httptransport.NewServer(
		endpoints.GetUserCard,
		commonHttp.MakeRequestDecoder(func(ctx context.Context, r *http.Request) (interface{}, error) {
			return &usersEndpoints.GetUserCardRequest{
				Id: middleware.ExtractUserId(ctx),
			}, nil
		}),
		commonHttp.MakeResponseEncoder(extendErrorStatusMap),
		options...,
	))

	return s
}
