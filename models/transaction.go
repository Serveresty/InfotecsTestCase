package models

import "time"

type Transaction struct {
	SendTime time.Time `json:"time"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	Amount   float32   `json:"amount"`
}
