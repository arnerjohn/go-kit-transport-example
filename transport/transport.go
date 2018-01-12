package transport

import (
	"context"
	"encoding/json"
	"github.com/arnerjohn/transport-example/service"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type EndpointSet struct {
	UppercaseEndpoint endpoint.Endpoint
	CountEndpoint     endpoint.Endpoint
}

func MakeEndpoints(svc service.ServiceInterface) EndpointSet {
	var uppercaseEndpoint endpoint.Endpoint
	uppercaseEndpoint = MakeUppercaseEndpoint(svc)

	var countEndpoint endpoint.Endpoint
	countEndpoint = MakeCountEndpoint(svc)

	return EndpointSet{
		UppercaseEndpoint: uppercaseEndpoint,
		CountEndpoint:     countEndpoint,
	}
}

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

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	return json.NewEncoder(w).Encode(response)
}
