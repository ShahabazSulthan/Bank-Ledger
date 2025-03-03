package repository

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ShahabazSulthan/BankingService/pkg/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name    string
		input   *domain.Account
		stub    func(sqlmock.Sqlmock)
		wantErr string
		wantID  uint64
	}{
		{
			name: "successfully inserted account",
			input: &domain.Account{
				Name:                 "Savings Account",
				PrimaryAccountNumber: "1234567890",
				Balance:              1000.0,
			},
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("INSERT INTO accounts").
					WithArgs("Savings Account", "1234567890", 1000.0).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			wantErr: "",
			wantID:  1,
		},
		{
			name: "error inserting account",
			input: &domain.Account{
				Name:                 "Current Account",
				PrimaryAccountNumber: "9876543210",
				Balance:              2000.0,
			},
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("INSERT INTO accounts").
					WithArgs("Current Account", "9876543210", 2000.0).
					WillReturnError(fmt.Errorf("failed to insert account"))
			},
			wantErr: "failed to insert account",
			wantID:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSql, _ := sqlmock.New()
			defer mockDB.Close()

			DB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			
			tt.stub(mockSql)

			accountRepo := &AccountRepo{DB: DB}

			accountID, err := accountRepo.CreateAccount(tt.input)

			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				assert.Equal(t, tt.wantID, accountID)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantID, accountID)
			}

			err = mockSql.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestGetAccount(t *testing.T) {
	tests := []struct {
		name    string
		id      uint64
		stub    func(sqlmock.Sqlmock)
		wantErr string
		want    *domain.Account
	}{
		{
			name: "successfully fetched account",
			id:   1,
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(`SELECT \* FROM "accounts" WHERE "accounts"\."id" = \$1 AND "accounts"\."deleted_at" IS NULL ORDER BY "accounts"\."id" LIMIT \$2`).
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "primary_account_number", "balance"}).AddRow(1, "Savings Account", "1234567890", 1000.0))
			},
			wantErr: "",
			want: &domain.Account{
				Model:                gorm.Model{ID: 1},
				Name:                 "Savings Account",
				PrimaryAccountNumber: "1234567890",
				Balance:              1000.0,
			},
		},
		{
			name: "account not found",
			id:   2,
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(`SELECT \* FROM "accounts" WHERE "accounts"\."id" = \$1 AND "accounts"\."deleted_at" IS NULL ORDER BY "accounts"\."id" LIMIT \$2`).
					WithArgs(2, 1).
					WillReturnError(gorm.ErrRecordNotFound)
			},
			wantErr: "",
			want:    nil,
		},
		{
			name: "failed to fetch account",
			id:   3,
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(`SELECT \* FROM "accounts" WHERE "accounts"\."id" = \$1 AND "accounts"\."deleted_at" IS NULL ORDER BY "accounts"\."id" LIMIT \$2`).
					WithArgs(3, 1).
					WillReturnError(fmt.Errorf("database error"))
			},
			wantErr: "failed to fetch account",
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		
			mockDB, mockSql, _ := sqlmock.New()
			defer mockDB.Close()

			DB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSql)

			accountRepo := &AccountRepo{DB: DB}


			account, err := accountRepo.GetAccount(tt.id)

			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, account)
			}

			err = mockSql.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
