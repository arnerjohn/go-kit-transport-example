package main

import (
	"github.com/arnerjohn/transport-example/service"
	"github.com/arnerjohn/transport-example/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

func main() {
	var svc service.ServiceInterface
	svc = service.Service{}

	uppercaseHandler := httptransport.NewServer(
		transport.MakeUppercaseEndpoint(svc),
		transport.DecodeUppercaseRequest,
		transport.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		transport.MakeCountEndpoint(svc),
		transport.DecodeCountRequest,
		transport.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
