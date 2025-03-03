package model

import "time"

type TransactionLog struct {
	TransactionID   string    `json:"transaction_id"`
	AccountID       string    `json:"account_id"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"` 
	Status          string    `json:"status"`          
	CreatedAt       time.Time `json:"created_at"`
}
