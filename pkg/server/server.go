package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/rMaxiQp/grpc-proto/greet"
)

type Server struct {
	greet.UnimplementedGreetServer
}

func (s *Server) SayHello(ctx context.Context, req *greet.HelloRequest) (*greet.HelloResponse, error) {
	prefix := "Hello"
	switch strings.ToLower(req.GetLanguage()) {
	case "es":
		prefix = "Hola"
	case "fr":
		prefix = "Bonjour"
	case "de":
		prefix = "Hallo"
	}
	return &greet.HelloResponse{Message: fmt.Sprintf("%s %s", prefix, req.GetName())}, nil
}
