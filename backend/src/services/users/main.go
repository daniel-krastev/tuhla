package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"tuhla/common/db"
	"tuhla/services/users/controller"
	"tuhla/services/users/dbstorage"
	"tuhla/services/users/proto/usersservicepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const connString = "postgresql://root@localhost:26257/tuhla?sslmode=disable"

func main() {
	ctx := context.Background()
	conn, err := db.Connect(ctx, connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	userController := controller.New(dbstorage.New(conn))

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 1123))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot listen on the localhost: %v\n", err)
		os.Exit(1)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	usersservicepb.RegisterUsersServer(grpcServer, userController)

	fmt.Println("Starting user service...")
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot serve grpc server: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("User service stopped...")
}
