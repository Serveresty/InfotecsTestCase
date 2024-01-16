package models

import "time"

type Transaction struct {
	SendTime time.Time
	From     string
	To       string  `json:"to"`
	Amount   float32 `json:"amount"`
}
