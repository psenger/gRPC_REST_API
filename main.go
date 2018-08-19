// generates the gRPC stubs
//go:generate protoc -I api -I$GOPATH/bin -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:api api/api.proto
// generates the rest-proxy-to-gRPC server
//go:generate protoc -I api -I$GOPATH/bin -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  --grpc-gateway_out=logtostderr=true:api api/api.proto

package main

import (
	"github.com/psenger/gRPC_REST_API/api"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"net"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	log "github.com/sirupsen/logrus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/psenger/gRPC_REST_API/server"
	"net/http"
	"fmt"
	"github.com/spf13/viper"
)

type AppServer struct {
	gRPCAddress string
	restAddress string
	server api.SimplServiceServer
}

func ( app *AppServer) startGRPCServer() error {
	lis, err := net.Listen("tcp", app.gRPCAddress )
	if err != nil {
		return err;
	} else {
		// https://github.com/grpc-ecosystem/go-grpc-middleware
		server := grpc.NewServer(
				grpc.UnaryInterceptor(
					grpc_middleware.ChainUnaryServer(
						grpc_logrus.UnaryServerInterceptor(
							log.NewEntry(
								log.StandardLogger(),
								),
							),
						),
					),
			)
		api.RegisterSimplServiceServer( server, app.server )
		reflection.Register(server)
		return server.Serve(lis)
	}
}

func ( app *AppServer) startRESTServer() error {
	ctx, cancel := context.WithCancel( context.Background() )
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{ grpc.WithInsecure() }

	err := api.RegisterSimplServiceHandlerFromEndpoint( ctx, mux, app.gRPCAddress, opts )
	if err != nil {
		return err
	} else {
		return http.ListenAndServe(  app.restAddress, mux  )
	}

}

func ( app *AppServer ) start() {
	errors := make(chan error)
	go func() { errors <- fmt.Errorf("GRPC Server Faield To Start: %v", app.startGRPCServer() )}()
	go func() { errors <- fmt.Errorf("REST Server Faield To Start: %v", app.startRESTServer() )}()
	log.Fatal(<-errors)
}

func init() {
	log.SetFormatter( &log.TextFormatter{FullTimestamp:true})
	log.SetLevel(log.DebugLevel)
}

func main() {
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	//err := viper.ReadInConfig() // Find and read the config file
	//if err != nil { // Handle errors reading the config file
	//	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	//}
	gRPCAddress := viper.GetString("gRPCAddress")
	fmt.Print(gRPCAddress)
	application := AppServer{
		gRPCAddress: ":9191",
		restAddress: ":8080",
		server:      server.NewGrpcHelloService(),
	}
	application.start()
}