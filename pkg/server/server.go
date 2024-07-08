package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/rMaxiQp/grpc-proto/proto/greet"
)

type Server struct{}

var _ greet.GreetServer = (*Server)(nil)

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := &greet.HelloRequest{
		Name:     r.URL.Query().Get("name"),
		Language: r.URL.Query().Get("language"),
	}
	resp, err := s.sayHello(r.Context(), body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := protojson.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	glog.Infof("result: %v", string(result))
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) sayHello(ctx context.Context, req *greet.HelloRequest) (*greet.HelloResponse, error) {
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

func (s *Server) SayHello(ctx context.Context, req *greet.HelloRequest) (*greet.HelloResponse, error) {
	return s.sayHello(ctx, req)
}
