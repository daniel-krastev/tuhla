package db

import (
	"context"
	"fmt"

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

	con, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB with [%s]: %w", connString, err)
	}

	return &DBConnection{con}, nil
}
