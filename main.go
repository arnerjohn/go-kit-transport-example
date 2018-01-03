package main

import (
	"context"
	"fmt"
	"github.com/arnerjohn/transport-example/service"
	"github.com/arnerjohn/transport-example/transport"
)

func main() {
	var svc service.ServiceInterface
	svc = service.Service{}

	uppercaseRequest := transport.UppercaseRequest{Input: "hello"}
	uppercaseEndpoint := transport.MakeUppercaseEndpoint(svc)
	uppercaseOutput, err := uppercaseEndpoint(context.Background(), uppercaseRequest)

	if err != nil {
		panic(err)
	}
	fmt.Println(uppercaseOutput)

	countRequest := transport.CountRequest{Input: "hello"}
	countEndpoint := transport.MakeCountEndpoint(svc)
	countOutput, err := countEndpoint(context.Background(), countRequest)

	if err != nil {
		panic(err)
	}
	fmt.Println(countOutput)
}
