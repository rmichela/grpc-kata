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

	// Call the C# server on port 9002
	fmt.Println("Press enter to call the C# server...")
	fmt.Scanln()

	// Call the service
	conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	req := HelloRequest{Name: "Golang"}

	resp, err := client.SayHello(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	for _, message := range resp.GetMessage() {
		fmt.Println(message)
	}

	// Wait to exit
	fmt.Scanln()
}
