package http

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

type RequestDtoFactory func() interface{}

type CustomDecoder func(ctx context.Context, r *http.Request) (interface{}, error)

func MakeRequestDecoder(decoder CustomDecoder) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		if decoder == nil {
			return nil, nil
		}

		req, err := decoder(ctx, r)
		if err != nil {
			return nil, err
		}

		return req, nil
	}
}

// todo delete
//func DecodeJsonBody(factory RequestDtoFactory) CustomDecoder {
//	return func(ctx context.Context, r *http.Request) (interface{}, error) {
//		var req interface{} = decodedValue
//
//		if req == nil {
//			req = factory()
//		}
//
//		if req == nil {
//			return nil, nil
//		}
//
//		var err error
//		converted, ok := req.(easyjson.Unmarshaler)
//		if ok {
//			err = easyjson.UnmarshalFromReader(r.Body, &converted)
//			if err == nil {
//				req = converted
//			}
//
//		} else {
//			err = json.NewDecoder(r.Body).Decode(&req)
//		}
//
//		if errors.Is(err, io.EOF) {
//			return nil, common.ErrInvalidJson
//		}
//
//		if err != nil {
//			return nil, err
//		}
//
//		return req, nil
//	}
//}
