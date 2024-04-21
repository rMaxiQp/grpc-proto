package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/rMaxiQp/grpc-proto/greet"
	"github.com/rMaxiQp/grpc-proto/pkg/server"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &server.Server{}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	greet.RegisterGreetServer(grpcServer, s)

	fmt.Println("Starting")
	grpcServer.Serve(lis)

	fmt.Println("Hello!")
}
