#!/usr/bin/env bash

mkdir src && cd src
mkdir github.com && cd github.com
mkdir psenger && cd psenger
mkdir gRPC_REST_API && cd gRPC_REST_API

mv ../../../../Gopkg.lock ./
mv ../../../../Gopkg.toml ./
mv ../../../../main.go ./
mv ../../../../README.md ./
mv ../../../../api ./
mv ../../../../server ./
cd ../../../../
export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH:$GOPATH/bin:$GOPATH/bin/include
go build -o main ./src/github.com/psenger/main.go