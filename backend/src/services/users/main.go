package main

import (
	"context"
	"fmt"
	"os"

	"tuhla/common/db"
	"tuhla/services/users/controller"
	_ "tuhla/services/users/proto/usersmodelpb"
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

	controller.New(conn)
}
