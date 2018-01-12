package main

import (
	"github.com/arnerjohn/transport-example/service"
	"github.com/arnerjohn/transport-example/transport"
	"log"
	"net"
	"net/http"
)

func main() {
	var svc service.ServiceInterface
	svc = service.Service{}

	var endpoints transport.EndpointSet
	endpoints = transport.MakeEndpoints(svc)

	var httpHandlers http.Handler
	httpHandlers = transport.NewHTTPHandler(endpoints)

	httpListener, err := net.Listen("tcp", ":8080")

	defer httpListener.Close()

	if err != nil {
		log.Fatal(err)
	}

	http.Serve(httpListener, httpHandlers)
}
