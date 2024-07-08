package main

import (
	"flag"
	"io"
	"net/http"

	"github.com/golang/glog"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/rMaxiQp/grpc-proto/proto/greet"
)

func main() {
	flag.Parse()
	client := &http.Client{}
	res, err := client.Get("http://localhost:9000/v1/greet/world?language=es&name=world")
	if err != nil {
		glog.Exitf("failed to request: %v", err)
	}
	defer res.Body.Close()
	bs, err := io.ReadAll(res.Body)
	var resp greet.HelloResponse
	if err := protojson.Unmarshal(bs, &resp); err != nil {
		glog.Exitf("failed to unmarshal: %v", err)
	}
	glog.Infof("response: %v", &resp)
}
