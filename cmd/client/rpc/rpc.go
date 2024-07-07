package main

import (
	"context"
	"flag"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	"github.com/rMaxiQp/grpc-proto/proto/greet"
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
	glog.V(2).Infof("resp message: %s", resp.Message)
}
