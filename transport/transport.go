package transport

import (
	"context"
	"github.com/arnerjohn/transport-example/service"
	"github.com/go-kit/kit/endpoint"
)

type UppercaseRequest struct {
	Input string `json:"input"`
}

type UppercaseResponse struct {
	Output string `json:"output"`
	Err    string `json:"err,omitempty"`
}

type CountRequest struct {
	Input string `json:"input"`
}

type CountResponse struct {
	Output int `json:"output"`
}

func MakeUppercaseEndpoint(svc service.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(ctx, req.Input)
		if err != nil {
			return UppercaseResponse{Output: v, Err: err.Error()}, nil
		}
		return UppercaseResponse{Output: v, Err: ""}, nil
	}
}

func MakeCountEndpoint(svc service.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(ctx, req.Input)
		return CountResponse{Output: v}, nil
	}
}
