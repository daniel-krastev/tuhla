package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type DBConnection struct {
	*pgx.Conn
}

type ConnectionConfig struct{}

func Connect(ctx context.Context, connString string) (*DBConnection, error) {
	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("error parsing DB connection string [%s]: %w", connString, err)
	}

	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB with [%s]: %w", connString, err)
	}

	// testConn(conn, ctx)
	return &DBConnection{conn}, nil
}

func testConn(conn *pgx.Conn, ctx context.Context) {
	conn.Exec(ctx, "USE tuhla")

	var name, address string
	err := conn.QueryRow(context.Background(), "select name, address from app.houses where id=$1", 3).Scan(&name, &address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, address)
}
