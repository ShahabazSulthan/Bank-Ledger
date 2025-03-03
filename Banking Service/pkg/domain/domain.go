package domain

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType string

const (
	Deposit    TransactionType = "deposit"
	Withdrawal TransactionType = "withdrawal"
)

type TransactionStatus string

const (
	Failed    TransactionStatus = "failed"
	Pending   TransactionStatus = "pending"
	Confirmed TransactionStatus = "confirmed"
)

type Account struct {
	gorm.Model
	Name                 string
	PrimaryAccountNumber string `gorm:"uniqueIndex"`
	Balance              float64
	Transactions         []Transaction `gorm:"foreignKey:AccountID"`
}

type Transaction struct {
	gorm.Model
	AccountID uint64          
	Account   Account         `gorm:"foreignKey:AccountID"` 
	Type      TransactionType `gorm:"type:varchar(20)"`     
	Amount    float64
	Status    TransactionStatus `gorm:"type:varchar(20);default:'pending'"` 
}

type TransactionLog struct {
	TransactionID   string    `json:"transaction_id"`
	AccountID       string    `json:"account_id"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"` 
	Status          string    `json:"status"`           
	CreatedAt       time.Time `json:"created_at"`
}
