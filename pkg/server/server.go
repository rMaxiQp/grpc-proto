package server

import "github.com/rMaxiQp/grpc-proto/greet"

func StartService() *greet.HelloResponse {
	return &greet.HelloResponse{Message: "Hello, World!"}
}
