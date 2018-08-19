package server

import (
	"github.com/psenger/gRPC_REST_API/api"
	"golang.org/x/net/context"
)

type GrpcHelloService struct {

}

func NewGrpcHelloService () api.SimplServiceServer {
	return &GrpcHelloService{}
}

func (s *GrpcHelloService) Echo(context.Context, *api.HelloRequest) (*api.HelloResponse, error) {
	var response api.HelloResponse
	response.Message = "hello there";
	return &response, nil
}