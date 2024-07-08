package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	"github.com/rMaxiQp/grpc-proto/pkg/server"
	"github.com/rMaxiQp/grpc-proto/proto/greet"
)

func main() {
	flag.Parse()
	s := &server.Server{}
	go func() {
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		greet.RegisterGreetServer(grpcServer, s)
		grpcPort := "9090"
		handler, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			glog.Exitf("failed to start grpc port: %w", err)
		}
		grpcServer.Serve(handler)
	}()
	port := "9000"
	if err := http.ListenAndServe(":"+port, s); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
