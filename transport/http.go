package transport

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func NewHTTPHandler(endpoints EndpointSet) http.Handler {
	m := http.NewServeMux()

	m.Handle("/uppercase", httptransport.NewServer(
		endpoints.UppercaseEndpoint,
		DecodeUppercaseRequest,
		EncodeResponse,
	))

	m.Handle("/count", httptransport.NewServer(
		endpoints.CountEndpoint,
		DecodeCountRequest,
		EncodeResponse,
	))

	return m
}
