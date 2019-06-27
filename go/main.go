package main

import (
	"context"
	"fmt"
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
	fmt.Println("Hello Golang")
}
