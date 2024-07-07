package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/rMaxiQp/grpc-proto/proto/greet"
)

func main() {
	flag.Parse()
	client := &http.Client{}
	res, err := client.Get("http://localhost:9000/v1/greet/world?lang=es")
	if err != nil {
		glog.Exitf("failed to request: %v", err)
	}
	defer res.Body.Close()
	var resp greet.HelloResponse
	var bs []byte
	if _, err := res.Body.Read(bs); err != nil {
		glog.Exitf("failed to read: %v", err)
	}
	if err := protojson.Unmarshal(bs, &resp); err != nil {
		glog.Exitf("failed to unmarshal: %v", err)
	}
	glog.Infof("response: %v", &resp)
}
