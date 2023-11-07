package main

import (
	"context"
	"flag"
	"fmt"
	"grpcdemo/greetingrpc"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", ":42069", "server adddress to connect to")
	name = flag.String("name", defaultName, "name to greet")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connection could not be established: %v\n", err)
	}
	defer conn.Close()

	c := greetingrpc.NewGreetingRpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetGreeting(ctx, &greetingrpc.Name{Name: *name})
	if err != nil {
		log.Fatalf("could not call greeting method: %v\n", err)
	}
	fmt.Printf("%s", r.GetGreeting())
}
