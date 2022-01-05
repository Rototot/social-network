package http

import (
	"context"
	"encoding/json"
	"net/http"
)
import httptransport "github.com/go-kit/kit/transport/http"

func MakeErrorEncoder(extendsStatusMapper ErrStatusMap) httptransport.ErrorEncoder {
	return func(_ context.Context, err error, w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(ErrorCodeStatusMapperFunc(err, extendsStatusMapper))

		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
	}
}
