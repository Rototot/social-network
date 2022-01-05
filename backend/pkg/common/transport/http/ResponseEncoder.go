package http

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func MakeResponseEncoder(extendsStatusMapper map[error]int) httptransport.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		if e, ok := response.(error); ok && e != nil {
			// Not a Go kit transport error, but a business-logic error.
			// Provide those as HTTP errors.
			handler := MakeErrorEncoder(extendsStatusMapper)
			handler(ctx, e, w)

			return nil
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(response)
	}
}
