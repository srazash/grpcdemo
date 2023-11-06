package main

import (
	"context"
	"fmt"
	"grpcdemo/greetingrpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myGreetingRpcServer struct {
	greetingrpc.UnimplementedGreetingRpcServer
}

func (m myGreetingRpcServer) GetGreeting(c context.Context, n *greetingrpc.Name) (*greetingrpc.Greeting, error) {
	return &greetingrpc.Greeting{Greeting: fmt.Sprintf("Hello, %s\n", n.String())}, nil
}

func main() {
	const netw string = "tcp"
	const addr string = ":42069"
	lis, err := net.Listen(netw, addr)

	fmt.Printf("server --> (%s) %s\n", netw, addr)

	if err != nil {
		log.Fatalf("unable to create listener: %s\n", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myGreetingRpcServer{}

	greetingrpc.RegisterGreetingRpcServer(serverRegistrar, service)
	if serverRegistrar.Serve(lis) != nil {
		log.Fatalf("unable to serve: %s\n", err)
	}
}
