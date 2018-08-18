// generates the gRPC stubs
//go:generate protoc -I api -i$GOPATH/bin/include -I$GOPATH/../../grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:api api/api.proto
// generates the rest-proxy-to-gRPC server
//go:generate protoc -I api -i$GOPATH/bin/include -I$GOPATH/../../grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway=plugins=logtostderr=true:api api/api.proto

package main

import (
	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
)

func main () {
	var opts []grpc.ServerOption
}
