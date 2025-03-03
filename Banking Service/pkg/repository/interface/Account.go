package interfaces

import "github.com/ShahabazSulthan/BankingService/pkg/domain"

type AccountRepository interface {
	CreateAccount(account *domain.Account) (uint64, error)
	GetAccount(id uint64) (*domain.Account, error)
	UpdateBalance(id uint64, amount float64) (float64, error)

	CreateTransaction(txn *domain.Transaction) (uint64, error)
	GetTransactionByID(id uint64) (*domain.Transaction, error)
	GetTransactionsByAccountID(accountID uint64) ([]domain.Transaction, error)
	UpdateTransactionStatus(id uint64, status domain.TransactionStatus) error
}
