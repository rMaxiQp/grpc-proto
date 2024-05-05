package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang/glog"

	"github.com/rMaxiQp/grpc-proto/proto/greet"
)

type Server struct{}

var _ greet.GreetServer = (*Server)(nil)

func (s *Server) SayHello(ctx context.Context, req *greet.HelloRequest) (*greet.HelloResponse, error) {
	glog.V(2).Infof("Received request: %v", req)
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
