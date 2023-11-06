package main

import (
	"context"
	"fmt"
	"grpcdemo/greetingrpc"

	"google.golang.org/grpc"
)

type myGreetingRpcClient struct{}

func (c *myGreetingRpcClient) GetGreeting(ctx context.Context, in *greetingrpc.Name, opts ...grpc.CallOption) (*greetingrpc.Greeting, error) {
	return nil, nil // implement me pls :)
}

func main() {
	const name string = "Ryan"
	fmt.Printf("%s lives!\n", name)
}
