package database

import (
	"InfotecsTestCase/cerr"
	"InfotecsTestCase/models"
	"context"
	"time"
)

func SendMoney(transaction *models.Transaction) error {
	row1 := DB.QueryRow(context.Background(), `SELECT wallet_id, balance FROM "wallets" WHERE wallet_id=$1`, transaction.From)

	var from, to models.Wallet
	err := row1.Scan(&from.Id, &from.Balance)
	if err != nil {
		return cerr.SenderWalletNotFound
	}

	row2 := DB.QueryRow(context.Background(), `SELECT wallet_id, balance FROM "wallets" WHERE wallet_id=$1`, transaction.To)
	err = row2.Scan(&to.Id, &to.Balance)
	if err != nil {
		return cerr.ReceiverWalletNotFound
	}

	if transaction.Amount > from.Balance {
		return cerr.LessMoney
	}

	from.Balance = from.Balance - transaction.Amount
	to.Balance = to.Balance + transaction.Amount

	_, err = DB.Exec(context.Background(), `UPDATE wallets SET balance=$1 WHERE wallet_id=$2`, from.Balance, from.Id)
	if err != nil {
		return err
	}

	_, err = DB.Exec(context.Background(), `UPDATE wallets SET balance=$1 WHERE wallet_id=$2`, to.Balance, to.Id)
	if err != nil {
		return err
	}

	_, err = DB.Exec(context.Background(), `INSERT INTO "transactions" (send_time, from_id, to_id, amount) VALUES($1, $2, $3, $4)`, time.Now(), from.Id, to.Id, transaction.Amount)
	if err != nil {
		return err
	}

	return nil
}
