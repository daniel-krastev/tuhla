package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

const connString = "postgresql://root@localhost:26257/defaultdb?sslmode=disable"

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	conn.Exec(ctx, "USE tuhla")
	var name, address string
	err = conn.QueryRow(context.Background(), "select name, address from app.houses where id=$1", 3).Scan(&name, &address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, address)
}
