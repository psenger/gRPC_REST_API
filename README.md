# README

cd /Users/psenger/Documents/Dev/gosample
export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH:$GOPATH/bin:$GOPATH/bin/include
echo $GOPATH
mkdir src && cd src
mkdir github.com && cd github.com
mkdir psenger && cd psenger
git clone git@github.com:psenger/gRPC_REST_API.git


cd $GOPATH
mkdir $GOPATH/bin
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

* go get this zip file with a browser.*
https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-osx-x86_64.zip
Copy all the code and directories into ```bin/*``` and ```$GOPATH/bin```
But the include should be in the bin directory

```
$ cd /Users/psenger/Downloads/

$ cd /Users/psenger/Downloads/protoc-3.6.1-osx-x86_64
$ mv bin/* $GOPATH/bin/
$ mv include/ $GOPATH/bin
$ cd $GOPATH/bin && ls -la
    total 38408
    .
    ..
    dep
    includea
    protoc
$

```

cd $GOPATH/src/github.com/psenger/gRPC_REST_API

* now make the go file *

---- main.go
// generates the gRPC stubs
//go:generate protoc -Iapi -I../../../../bin/include -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:api api/api.proto
// generates the rest-proxy-to-gRPC server
//go:generate protoc -Iapi -I../../../../bin/include -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway=plugins=logtostderr=true:api api/api.proto

package main

import (
	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
)

func main () {
	var opts []grpc.ServerOption
}



--- api/api.proto

syntax = "proto3";

package api;

import "google/api/annotations.proto";

message HelloResponse {
    string message = 1;
}

message HelloRequest {
    string name = 1;
}

service SimplService {
    rpc Echo (HelloRequest) returns (HelloResponse) {
        option (google.api.http) = {
            get:"/v1/echo"
        };
    }
}
---

cd $GOPATH/src/github.com/psenger/gRPC_REST_API
dep init



cd $GOPATH/src/github.com/psenger/gRPC_REST_API

go get -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go

dep ensure -add github.com/golang/protobuf/proto
dep ensure -add github.com/golang/protobuf/protoc-gen-go
dep ensure -add google.golang.org/grpcc
dep ensure -add github.com/golang/protobuf <<--- this doesnt install?
dep ensure -add github.com/grpc-ecosystem/grpc-gateway
dep ensure -add github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
dep ensure -add github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
