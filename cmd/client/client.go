package main

import (
	"context"
	"flag"
	"fmt"

	"google.golang.org/grpc"

	"github.com/rMaxiQp/grpc-proto/greet"
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(":9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := greet.NewGreetClient(conn)
	resp, err := client.SayHello(context.Background(), &greet.HelloRequest{Name: "World"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp message: %s\n", resp.Message)
}
