package ping

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	commonHttp "social-network/pkg/common/transport/http"
)

func MakePingHttpHandler(
	endpoints Endpoints,
	logger log.Logger,
) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(commonHttp.MakeErrorEncoder(nil)),
	}

	r.Methods("GET", "OPTIONS", "HEAD").Path("/").Handler(httptransport.NewServer(
		endpoints.Ping,
		commonHttp.MakeRequestDecoder(nil),
		commonHttp.MakeResponseEncoder(nil),
		options...,
	))

	return r
}
