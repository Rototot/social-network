package http

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

type RequestDtoFactory func() any

func MakeRequestDecoder(factory RequestDtoFactory) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (request interface{}, err error) {
		var req = factory()
		if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
			return nil, e
		}

		return req, nil
	}
}
