package http

import (
	"context"
	"encoding/json"
	"net/http"
)
import httptransport "github.com/go-kit/kit/transport/http"

func MakeErrorEncoder() httptransport.ErrorEncoder {
	return func(_ context.Context, err error, w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(ErrorCodeStatusMapperFunc(err, map[error]int{}))

		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
	}
}
