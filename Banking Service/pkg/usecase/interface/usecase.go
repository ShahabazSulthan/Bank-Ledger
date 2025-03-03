package interface_usecase

import "github.com/ShahabazSulthan/BankingService/pkg/domain"

type AccountUsecase interface {
	CreateAccount(account *domain.Account) (uint64, error)
	GetAccount(id uint64) (*domain.Account, error)
	UpdateBalance(id uint64, amount float64) (float64, error)

	ProcessTransaction(txn *domain.Transaction) (uint64, error)
	GetTransaction(id uint64) (*domain.Transaction, error)
	GetTransactionsByAccount(accountID uint64) ([]domain.Transaction, error)
}
