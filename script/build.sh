
cd ..
export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH:$GOPATH/bin:$GOPATH/bin/include

mkdir src && cd src
mkdir github.com && cd github.com
mkdir psenger && cd psenger
mkdir gRPC_REST_API && cd gRPC_REST_API

cd $GOPATH

mv Gopkg.lock ./src/github.com/psenger/gRPC_REST_API/
mv Gopkg.toml ./src/github.com/psenger/gRPC_REST_API/
mv main.go    ./src/github.com/psenger/gRPC_REST_API/
mv README.md  ./src/github.com/psenger/gRPC_REST_API/
mv api        ./src/github.com/psenger/gRPC_REST_API/
mv server     ./src/github.com/psenger/gRPC_REST_API/
mv vendor     ./src/github.com/psenger/gRPC_REST_API/

go build -o main ./src/github.com/psenger/gRPC_REST_API/main.go
