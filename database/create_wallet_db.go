package database

import (
	"InfotecsTestCase/models"
	"context"
)

func CreateWalletDB(wallet *models.Wallet) error {
	_, err := DB.Exec(context.Background(), `INSERT INTO "wallets" (wallet_id, balance) VALUES($1, $2)`, wallet.Id, wallet.Balance)
	if err != nil {
		return err
	}
	return nil
}
