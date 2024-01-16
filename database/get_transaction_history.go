package database

import (
	"InfotecsTestCase/cerr"
	"InfotecsTestCase/models"
	"context"
)

func GetTransactionHistory(id string) ([]models.Transaction, error) {
	row := DB.QueryRow(context.Background(), `SELECT balance FROM "wallets" WHERE wallet_id=$1`, id)

	var balance float32
	err := row.Scan(&balance)
	if err != nil {
		return nil, cerr.WalletNotFound
	}

	transactions, err := DB.Query(context.Background(), `SELECT send_time, from_id, to_id, amount FROM "transactions" WHERE from_id=$1 or to_id=$2`, id, id)
	if err != nil {
		return nil, cerr.WalletNotFound
	}

	var trns []models.Transaction
	for transactions.Next() {
		var trn models.Transaction
		err := transactions.Scan(&trn.SendTime, &trn.From, &trn.To, &trn.Amount)
		if err != nil {
			return nil, err
		}
		trns = append(trns, trn)
	}

	return trns, nil
}
