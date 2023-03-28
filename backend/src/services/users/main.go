package main

import (
	// "context"
	"flag"
	"fmt"
	"net"
	"os"

	_ "tuhla/common/db"
	"tuhla/services/users/controller"
	_ "tuhla/services/users/dbstorage"
	"tuhla/services/users/proto/usersservicepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	databaseURL    = flag.String("database_url", "postgresql://root@localhost:26257/tuhla?sslmode=disable", "database connection URL")
	serviceAddress = flag.String("service_address", "", "service address")
	servicePort    = flag.Int("service_port", 1124, "service port")
)

// force docker rebuild 12

func main() {
	flag.Parse()

	// ctx = context.Background()

	// conn, err := db.Connect(ctx, *databaseURL)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(ctx)

	// controller := controller.New(dbstorage.New(conn))
	controller := controller.New(nil)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *serviceAddress, *servicePort))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot listen on the localhost: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	reflection.Register(grpcServer)
	usersservicepb.RegisterUsersServer(grpcServer, controller)

	fmt.Println("Starting user service...")
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot serve grpc server: %v\n", err)
		os.Exit(1)
	}
}
