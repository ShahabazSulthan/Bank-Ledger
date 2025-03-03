package server

import (
	"context"

	"github.com/ShahabazSulthan/BankingService/pkg/domain"
	"github.com/ShahabazSulthan/BankingService/pkg/pb"
	interface_usecase "github.com/ShahabazSulthan/BankingService/pkg/usecase/interface"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
	AccountUsecase interface_usecase.AccountUsecase
}

func NewAccountServer(usecase interface_usecase.AccountUsecase) *AccountServer {
	return &AccountServer{AccountUsecase: usecase}
}

func (a *AccountServer) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	if req.Name == "" || req.PrimaryAccountNumber == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid account details")
	}

	account := &domain.Account{
		Name:                 req.Name,
		PrimaryAccountNumber: req.PrimaryAccountNumber,
		Balance:              req.Balance,
	}

	accountID, err := a.AccountUsecase.CreateAccount(account)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account: %v", err)
	}

	return &pb.CreateAccountResponse{
		AccountId: accountID,
		Message:   "Account created successfully",
	}, nil
}

func (a *AccountServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid account ID")
	}

	account, err := a.AccountUsecase.GetAccount(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "account not found: %v", err)
	}

	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:                   uint64(account.ID),
			Name:                 account.Name,
			PrimaryAccountNumber: account.PrimaryAccountNumber,
			Balance:              account.Balance,
		},
	}, nil
}

func (a *AccountServer) UpdateBalance(ctx context.Context, req *pb.UpdateBalanceRequest) (*pb.UpdateBalanceResponse, error) {
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "account ID is required")
	}
	if req.Amount == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "amount must be greater than zero")
	}

	updatedBalance, err := a.AccountUsecase.UpdateBalance(req.Id, req.Amount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update balance: %v", err)
	}

	return &pb.UpdateBalanceResponse{
		UpdatedBalance: updatedBalance,
		Message:        "Balance updated successfully",
	}, nil
}

func (a *AccountServer) ProcessTransaction(ctx context.Context, req *pb.ProcessTransactionRequest) (*pb.ProcessTransactionResponse, error) {
	if req.AccountId == 0 || req.Amount == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid transaction request")
	}

	txn := &domain.Transaction{
		AccountID: req.AccountId,
		Type:      convertProtoTypeToDomain(req.Type),
		Amount:    req.Amount,
		Status:    domain.Pending,
	}

	transactionID, err := a.AccountUsecase.ProcessTransaction(txn)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to process transaction: %v", err)
	}

	return &pb.ProcessTransactionResponse{
		TransactionId: transactionID,
		Message:       "Transaction processed successfully",
	}, nil
}

func (a *AccountServer) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid transaction ID")
	}

	txn, err := a.AccountUsecase.GetTransaction(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "transaction not found: %v", err)
	}

	return &pb.GetTransactionResponse{
		Transaction: &pb.Transaction{
			Id:        uint64(txn.ID),
			AccountId: txn.AccountID,
			Type:      convertDomainTypeToProto(txn.Type),
			Amount:    txn.Amount,
			Status:    convertDomainStatusToProto(txn.Status),
		},
	}, nil
}

func (a *AccountServer) GetTransactionsByAccount(ctx context.Context, req *pb.GetTransactionsByAccountRequest) (*pb.GetTransactionsByAccountResponse, error) {
	if req.AccountId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid account ID")
	}

	txns, err := a.AccountUsecase.GetTransactionsByAccount(req.AccountId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch transactions: %v", err)
	}

	pbTxns := make([]*pb.Transaction, 0, len(txns))
	for _, txn := range txns {
		pbTxns = append(pbTxns, &pb.Transaction{
			Id:        uint64(txn.ID),
			AccountId: txn.AccountID,
			Type:      convertDomainTypeToProto(txn.Type),
			Amount:    txn.Amount,
			Status:    convertDomainStatusToProto(txn.Status),
		})
	}

	return &pb.GetTransactionsByAccountResponse{Transactions: pbTxns}, nil
}

func convertProtoTypeToDomain(protoType pb.TransactionType) domain.TransactionType {
	switch protoType {
	case pb.TransactionType_DEPOSIT:
		return domain.Deposit
	case pb.TransactionType_WITHDRAWAL:
		return domain.Withdrawal
	default:
		return ""
	}
}

func convertDomainTypeToProto(domainType domain.TransactionType) pb.TransactionType {
	switch domainType {
	case domain.Deposit:
		return pb.TransactionType_DEPOSIT
	case domain.Withdrawal:
		return pb.TransactionType_WITHDRAWAL
	default:
		return pb.TransactionType_DEPOSIT
	}
}

func convertDomainStatusToProto(domainStatus domain.TransactionStatus) pb.TransactionStatus {
	switch domainStatus {
	case domain.Pending:
		return pb.TransactionStatus_PENDING
	case domain.Confirmed:
		return pb.TransactionStatus_CONFIRMED
	case domain.Failed:
		return pb.TransactionStatus_FAILED
	default:
		return pb.TransactionStatus_PENDING
	}
}
