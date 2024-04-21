package main

import (
	"log"

	"github.com/rMaxiQp/grpc-proto/pkg/server"
)

func main() {
	r := server.StartService()
	log.Println("%v", r)
}
