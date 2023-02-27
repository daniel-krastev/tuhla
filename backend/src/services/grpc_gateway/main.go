package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"tuhla/services/grpc_gateway/controller"
	"tuhla/services/grpc_gateway/proto/houses/housespb"
	"tuhla/services/grpc_gateway/proto/users/userspb"
	"tuhla/services/houses/proto/housesservicepb"
	"tuhla/services/users/proto/usersservicepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var (
	serviceAddress = flag.String("service_address", "localhost", "service address")
	servicePort    = flag.Int("service_port", 1123, "service port")

	// users address
	usersAddress = flag.String("users_address", "localhost", "users address")
	usersPort    = flag.Int("users_port", 1124, "users port")

	// houses address
	housesAddress = flag.String("houses_address", "localhost", "houses address")
	housesPort    = flag.Int("houses_port", 1125, "houses port")
)

func main() {
	flag.Parse()

	// Create internal service clients.
	usersConn, err := grpc.Dial(fmt.Sprintf("%s:%d", *usersAddress, *usersPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect to users service: %v", err)
	}
	usersClient := usersservicepb.NewUsersClient(usersConn)

	housesConn, err := grpc.Dial(fmt.Sprintf("%s:%d", *housesAddress, *housesPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect to houses service: %v", err)
	}
	housesClient := housesservicepb.NewHousesClient(housesConn)

	controller := controller.New(usersClient, housesClient)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *serviceAddress, *servicePort))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot listen on the localhost: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	reflection.Register(grpcServer)
	userspb.RegisterUsersServer(grpcServer, controller)
	housespb.RegisterHousesServer(grpcServer, controller)

	fmt.Println("Starting user service...")
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot serve grpc server: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("User service stopped...")
}
