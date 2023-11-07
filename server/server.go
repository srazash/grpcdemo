package main

import (
	"context"
	"fmt"
	"grpcdemo/greetingrpc"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type myGreetingRpcServer struct {
	greetingrpc.UnimplementedGreetingRpcServer
}

func (s myGreetingRpcServer) GetGreeting(ctx context.Context, name *greetingrpc.Name) (*greetingrpc.Greeting, error) {
	str := strings.Split(name.String(), "\"")[1]
	return &greetingrpc.Greeting{Greeting: fmt.Sprintf("Hello, %s!\n", str)}, nil
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
