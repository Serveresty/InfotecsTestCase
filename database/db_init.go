package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func DBConnectionInit(dbUrl string) error {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return fmt.Errorf("Error while connecting to database: %v", err)
	}

	DB = conn
	return nil
}
