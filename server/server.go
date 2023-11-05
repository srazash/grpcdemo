package main

import (
	"context"
	"fmt"
	"grpcdemo/server/rpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myRpcServer struct {
}

func (s myRpcServer) GetGreeting(c context.Context, n *rpc.Name) (*rpc.Greeting, error) {
	return &rpc.Greeting{Greeting: fmt.Sprintf("Hello, %s\n", n.GetName())}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":42069")

	if err != nil {
		log.Fatalf("Unable to create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myRpcServer{}
	rpc.RegisterRpcServer(serverRegistrar, service)
}
