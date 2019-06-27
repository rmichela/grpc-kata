package main

import (
	"context"
	"fmt"
	"net"

	grpc "google.golang.org/grpc"
)

type greeter struct{}

func (g *greeter) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	fmt.Println("Go service request: " + req.GetName())
	resp := HelloResponse{
		Message: []string{
			"Hello " + req.GetName(),
			"Aloha " + req.GetName(),
			"Howdy " + req.GetName()},
	}
	return &resp, nil
}

func main() {
	// Build the server
	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	RegisterGreeterServer(server, &greeter{})

	go func() {
		fmt.Println("Starting Go server on port 9001")
		err = server.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()
}
