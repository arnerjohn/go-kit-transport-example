package main

import (
	"context"
	"fmt"
	"github.com/arnerjohn/transport-example/service"
)

func main() {
	var svc service.ServiceInterface
	svc = service.Service{}

	upperOut, err := svc.Uppercase(context.Background(), "hello")

	if err != nil {
		panic(err)
	}

	fmt.Println(upperOut)

	countOut := svc.Count(context.Background(), "hello")

	fmt.Println(countOut)
}
