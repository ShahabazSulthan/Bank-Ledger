package repository

import (
	"errors"
	"fmt"

	"github.com/ShahabazSulthan/BankingService/pkg/domain"
	interfaces "github.com/ShahabazSulthan/BankingService/pkg/repository/interface"
	"gorm.io/gorm"
)

type AccountRepo struct {
	DB *gorm.DB
}

func NewAccountRepo(db *gorm.DB) interfaces.AccountRepository {
	return &AccountRepo{DB: db}
}

func (a *AccountRepo) CreateAccount(account *domain.Account) (uint64, error) {
	query := "INSERT INTO accounts (name, primary_account_number, balance, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW()) RETURNING id"
	var accountID uint64
	err := a.DB.Raw(query, account.Name, account.PrimaryAccountNumber, account.Balance).Scan(&accountID).Error
	if err != nil {
		return 0, fmt.Errorf("failed to insert account: %w", err)
	}
	return accountID, nil
}

func (r *AccountRepo) GetAccount(id uint64) (*domain.Account, error) {
	var account domain.Account
	if err := r.DB.First(&account, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch account: %w", err)
	}
	return &account, nil
}

func (a *AccountRepo) UpdateBalance(id uint64, amount float64) (float64, error) {
	result := a.DB.Model(&domain.Account{}).Where("id = ?", id).Update("balance", gorm.Expr("balance + ?", amount))

	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected == 0 {
		return 0, errors.New("no account updated, possible invalid ID")
	}

	var updatedAccount domain.Account
	if err := a.DB.First(&updatedAccount, id).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch updated balance: %w", err)
	}

	return updatedAccount.Balance, nil
}

func (a *AccountRepo) CreateTransaction(txn *domain.Transaction) (uint64, error) {
	if err := a.DB.Create(txn).Error; err != nil {
		return 0, fmt.Errorf("failed to insert transaction: %w", err)
	}
	return uint64(txn.ID), nil
}

func (a *AccountRepo) GetTransactionByID(id uint64) (*domain.Transaction, error) {
	var txn domain.Transaction
	if err := a.DB.First(&txn, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}
		return nil, err
	}
	return &txn, nil
}

func (a *AccountRepo) GetTransactionsByAccountID(accountID uint64) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := a.DB.Where("account_id = ?", accountID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (a *AccountRepo) UpdateTransactionStatus(id uint64, status domain.TransactionStatus) error {
	result := a.DB.Model(&domain.Transaction{}).Where("id = ?", id).Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no transaction updated, possible invalid ID")
	}

	return nil
}
