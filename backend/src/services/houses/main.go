package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"tuhla/services/houses/controller"
	"tuhla/services/houses/proto/housesservicepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	_              = flag.String("database_url", "postgresql://root@localhost:26257/tuhla?sslmode=disable", "database connection URL")
	serviceAddress = flag.String("service_address", "localhost", "service address")
	servicePort    = flag.Int("service_port", 1125, "service port")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *serviceAddress, *servicePort))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot listen on the localhost: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)

	reflection.Register(grpcServer)

	controller := controller.New()
	housesservicepb.RegisterHousesServer(grpcServer, controller)

	fmt.Println("Starting houses service...")
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot serve grpc server: %v\n", err)
		os.Exit(1)
	}
}
