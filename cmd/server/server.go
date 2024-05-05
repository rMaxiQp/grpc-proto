package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/rMaxiQp/grpc-proto/pkg/server"
	"github.com/rMaxiQp/grpc-proto/proto/greet"
)

func main() {
	flag.Parse()
	handler, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &server.Server{}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	greet.RegisterGreetServer(grpcServer, s)

	grpcServer.Serve(handler)
}
