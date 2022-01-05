package ping

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Ping endpoint.Endpoint
}

func MakeEndpoints() Endpoints {
	return Endpoints{
		Ping: MakePingEndpoint(),
	}
}

func MakePingEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return "ok", nil
	}
}
