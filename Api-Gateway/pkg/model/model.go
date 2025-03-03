package model

import "time"

type Account struct {
	ID                   uint64    `json:"id"`
	Name                 string    `json:"name"`
	PrimaryAccountNumber string    `json:"primary_account_number"`
	Balance              float64   `json:"balance"`
	CreatedAt            time.Time `json:"created_at"`
}

type Transaction struct {
	ID        uint64    `json:"id"`
	AccountID uint64    `json:"account_id"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type CommonResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      *string     `json:"error,omitempty"`
}

const (
	Deposit    = "deposit"
	Withdrawal = "withdrawal"
)

const (
	Pending   = "pending"
	Confirmed = "confirmed"
	Failed    = "failed"
)

type TransactionLog struct {
	TransactionID   string  `json:"transaction_id"`
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	Status          string  `json:"status"`
	CreatedAt       string  `json:"created_at"`
}
