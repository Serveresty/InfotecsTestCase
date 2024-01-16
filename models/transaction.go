package models

type Transaction struct {
	From   string
	To     string  `json:"to"`
	Amount float32 `json:"amount"`
}
