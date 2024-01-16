package database

import (
	"InfotecsTestCase/models"
	"context"
)

func GetWallet(id string) (*models.Wallet, error) {
	row := DB.QueryRow(context.Background(), `SELECT balance FROM "wallets" WHERE wallet_id=$1`, id)

	var balance float32
	err := row.Scan(&balance)
	if err != nil {
		return &models.Wallet{}, err
	}

	var wallet models.Wallet
	wallet.Id = id
	wallet.Balance = balance

	return &wallet, nil
}
