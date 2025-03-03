package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/ShahabazSulthan/BankingService/pkg/domain"
	interfaces "github.com/ShahabazSulthan/BankingService/pkg/repository/interface"
	interface_usecase "github.com/ShahabazSulthan/BankingService/pkg/usecase/interface"
	interface_kafka "github.com/ShahabazSulthan/BankingService/pkg/utils/kafka/interfaces"
)

type AccountUseCase struct {
	Repo  interfaces.AccountRepository
	Kafka interface_kafka.IKafkaProducer
}

func NewAccountUseCase(repo interfaces.AccountRepository, kafka interface_kafka.IKafkaProducer) interface_usecase.AccountUsecase {
	return &AccountUseCase{Repo: repo, Kafka: kafka}
}

func (u *AccountUseCase) CreateAccount(account *domain.Account) (uint64, error) {
	existingAccount, err := u.Repo.GetAccount(uint64(account.ID))
	if err != nil {
		return 0, fmt.Errorf("failed to check existing account: %w", err)
	}
	if existingAccount != nil {
		return 0, errors.New("account already exists")
	}

	accountID, err := u.Repo.CreateAccount(account)
	if err != nil {
		return 0, fmt.Errorf("failed to create account: %w", err)
	}
	return accountID, nil
}

func (u *AccountUseCase) GetAccount(id uint64) (*domain.Account, error) {
	account, err := u.Repo.GetAccount(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve account: %w", err)
	}
	return account, nil
}

func (u *AccountUseCase) UpdateBalance(id uint64, amount float64) (float64, error) {
	updatedBalance, err := u.Repo.UpdateBalance(id, amount)
	if err != nil {
		return 0, fmt.Errorf("failed to update balance: %w", err)
	}
	return updatedBalance, nil
}

func (u *AccountUseCase) ProcessTransaction(txn *domain.Transaction) (uint64, error) {
	account, err := u.Repo.GetAccount(txn.AccountID)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch account: %w", err)
	}
	if account == nil {
		return 0, errors.New("account does not exist")
	}

	if txn.Type == domain.Withdrawal {
		if account.Balance < txn.Amount {
			return 0, errors.New("insufficient funds")
		}
		txn.Amount = -txn.Amount
	} else if txn.Type != domain.Deposit {
		return 0, errors.New("invalid transaction type")
	}

	if _, err := u.Repo.UpdateBalance(txn.AccountID, txn.Amount); err != nil {
		return 0, fmt.Errorf("failed to update account balance: %w", err)
	}

	txn.Status = domain.Confirmed
	transactionID, err := u.Repo.CreateTransaction(txn)
	if err != nil {
		return 0, fmt.Errorf("failed to create transaction record: %w", err)
	}

	transactionLog := domain.TransactionLog{
		TransactionID:   fmt.Sprintf("%d", transactionID),
		AccountID:       fmt.Sprintf("%d", txn.AccountID),
		Amount:          txn.Amount,
		TransactionType: string(txn.Type),
		Status:          string(txn.Status),
		CreatedAt:       time.Now(),
	}

	if err := u.Kafka.KafkaNotificationProducer(&transactionLog); err != nil {
		return 0, fmt.Errorf("failed to publish transaction log to Kafka: %w", err)
	}

	return transactionID, nil
}

func (u *AccountUseCase) GetTransaction(id uint64) (*domain.Transaction, error) {
	txn, err := u.Repo.GetTransactionByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction: %w", err)
	}
	return txn, nil
}

func (u *AccountUseCase) GetTransactionsByAccount(accountID uint64) ([]domain.Transaction, error) {
	txns, err := u.Repo.GetTransactionsByAccountID(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions for account %d: %w", accountID, err)
	}
	return txns, nil
}
