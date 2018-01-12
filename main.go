package main

import (
	"context"
	"fmt"
	"github.com/arnerjohn/transport-example/pb"
	"github.com/arnerjohn/transport-example/service"
	"github.com/arnerjohn/transport-example/transport"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

func main() {
	app := cli.NewApp()
	app.Name = "Transport HTTP and gRPC Example"
	app.Commands = []cli.Command{
		{
			Name:   "listen",
			Action: listen,
		},
		{
			Name:   "client",
			Action: client,
		},
	}

	runError := app.Run(os.Args)
	if runError != nil {
		panic(runError)
	}
}

func listen(_ *cli.Context) {
	var svc service.ServiceInterface
	svc = service.Service{}

	var endpoints transport.EndpointSet
	endpoints = transport.MakeEndpoints(svc)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		var httpHandlers http.Handler
		httpHandlers = transport.NewHTTPHandler(endpoints)

		httpListener, err := net.Listen("tcp", ":8080")

		defer httpListener.Close()

		if err != nil {
			log.Fatal(err)
		}

		http.Serve(httpListener, httpHandlers)
	}()

	go func() {
		defer wg.Done()

		grpcListener, err := net.Listen("tcp", ":50051")

		defer grpcListener.Close()

		if err != nil {
			log.Fatal(err)
		}

		var grpcServer pb.TransportExampleServer
		grpcServer = transport.NewGRPCServer(endpoints)

		baseGRPCServer := grpc.NewServer()
		pb.RegisterTransportExampleServer(baseGRPCServer, grpcServer)
		baseGRPCServer.Serve(grpcListener)
	}()

	wg.Wait()
}

func client(cli *cli.Context) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTransportExampleClient(conn)

	var r interface{}
	switch cli.Args().First() {
	case "uppercase":
		r, err = client.Uppercase(context.Background(), &pb.Request{Input: cli.Args().Get(1)})
	case "count":
		r, err = client.Count(context.Background(), &pb.Request{Input: cli.Args().Get(1)})
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}
