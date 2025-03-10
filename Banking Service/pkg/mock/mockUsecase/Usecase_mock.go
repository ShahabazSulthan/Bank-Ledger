// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interface/usecase.go

// Package mockUsecase is a generated GoMock package.
package mockUsecase

import (
	reflect "reflect"

	domain "github.com/ShahabazSulthan/BankingService/pkg/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountUsecase is a mock of AccountUsecase interface.
type MockAccountUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAccountUsecaseMockRecorder
}

// MockAccountUsecaseMockRecorder is the mock recorder for MockAccountUsecase.
type MockAccountUsecaseMockRecorder struct {
	mock *MockAccountUsecase
}

// NewMockAccountUsecase creates a new mock instance.
func NewMockAccountUsecase(ctrl *gomock.Controller) *MockAccountUsecase {
	mock := &MockAccountUsecase{ctrl: ctrl}
	mock.recorder = &MockAccountUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountUsecase) EXPECT() *MockAccountUsecaseMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAccountUsecase) CreateAccount(account *domain.Account) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", account)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountUsecaseMockRecorder) CreateAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountUsecase)(nil).CreateAccount), account)
}

// GetAccount mocks base method.
func (m *MockAccountUsecase) GetAccount(id uint64) (*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", id)
	ret0, _ := ret[0].(*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountUsecaseMockRecorder) GetAccount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountUsecase)(nil).GetAccount), id)
}

// GetTransaction mocks base method.
func (m *MockAccountUsecase) GetTransaction(id uint64) (*domain.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", id)
	ret0, _ := ret[0].(*domain.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockAccountUsecaseMockRecorder) GetTransaction(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockAccountUsecase)(nil).GetTransaction), id)
}

// GetTransactionsByAccount mocks base method.
func (m *MockAccountUsecase) GetTransactionsByAccount(accountID uint64) ([]domain.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsByAccount", accountID)
	ret0, _ := ret[0].([]domain.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsByAccount indicates an expected call of GetTransactionsByAccount.
func (mr *MockAccountUsecaseMockRecorder) GetTransactionsByAccount(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsByAccount", reflect.TypeOf((*MockAccountUsecase)(nil).GetTransactionsByAccount), accountID)
}

// ProcessTransaction mocks base method.
func (m *MockAccountUsecase) ProcessTransaction(txn *domain.Transaction) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTransaction", txn)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessTransaction indicates an expected call of ProcessTransaction.
func (mr *MockAccountUsecaseMockRecorder) ProcessTransaction(txn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTransaction", reflect.TypeOf((*MockAccountUsecase)(nil).ProcessTransaction), txn)
}

// UpdateBalance mocks base method.
func (m *MockAccountUsecase) UpdateBalance(id uint64, amount float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalance", id, amount)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBalance indicates an expected call of UpdateBalance.
func (mr *MockAccountUsecaseMockRecorder) UpdateBalance(id, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockAccountUsecase)(nil).UpdateBalance), id, amount)
}
