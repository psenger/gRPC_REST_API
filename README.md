# README

## Steps I did to create this repo suitable to download

```bash
cd /Users/psenger/Documents/Dev/gosample
export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH:$GOPATH/bin:$GOPATH/bin/include
echo $GOPATH
mkdir src && cd src
mkdir github.com && cd github.com
mkdir psenger && cd psenger
git clone git@github.com:psenger/gRPC_REST_API.git
```

_Enter password from *git clone*_

```
cd $GOPATH
mkdir $GOPATH/bin
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

* go get this zip file with a browser.*
1. [https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-osx-x86_64.zip](https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-osx-x86_64.zip)
2. Copy everything _protoc-3.6.1-osx-x86_64/bin/_ from the zip file into ```$GOPATH/bin/``` directory ( so you will move this down a directory from where it was stored in the zip file ).
3. Copy everything _protoc-3.6.1-osx-x86_64/lib/_ from the zip file into ```$GOPATH/bin/lib``` directory.

```
cd /Users/psenger/Downloads/
rm -rf protoc-3.6.1-osx-x86_64
unzip protoc-3.6.1-osx-x86_64.zip -d protoc-3.6.1-osx-x86_64
cd /Users/psenger/Downloads/protoc-3.6.1-osx-x86_64
mv bin/* $GOPATH/bin/
mv include/ $GOPATH/bin
```



* now make the go file *

```
cd $GOPATH/src/github.com/psenger/gRPC_REST_API
```

main.go
----
```
cat > main.go
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

```

api/api.prot
----
```
cat > api/api.proto
 

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
```

----

```
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
```

Now check it in to github, the vendor files should go as well as all the Gopkg stuff.

## Then when you check out

```
export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH:$GOPATH/bin:$GOPATH/bin/include
ln -sf vendor src
go build main.go
```
