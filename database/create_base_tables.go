package database

import (
	"context"
	"fmt"
)

func CreateBaseTables() error {
	_, err := DB.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS "wallets" (wallet_id VARCHAR(50) UNIQUE, balance REAL NOT NULL);
		CREATE TABLE IF NOT EXISTS "transactions" (transaction_id bigserial UNIQUE, send_time TIMESTAMP NOT NULL, from_id VARCHAR(50) references wallets (wallet_id) NOT NULL, to_id VARCHAR(50) references wallets (wallet_id) NOT NULL, amount REAL NOT NULL);
		`)
	if err != nil {
		return fmt.Errorf("Error while creating base tables: %v", err)
	}

	return nil
}
