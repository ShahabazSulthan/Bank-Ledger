package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ShahabazSulthan/Api-Gateway/pkg/pb"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	Client pb.AccountServiceClient
}

func NewAccountHandler(client pb.AccountServiceClient) *AccountHandler {
	return &AccountHandler{Client: client}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req pb.CreateAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	resp, err := h.Client.CreateAccount(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp.Message})
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	resp, err := h.Client.GetAccount(context.Background(), &pb.GetAccountRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": resp.Account})
}

func (h *AccountHandler) UpdateBalance(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID is required"})
		return
	}

	accountID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID format"})
		return
	}

	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	resp, err := h.Client.UpdateBalance(context.Background(), &pb.UpdateBalanceRequest{
		Id:     accountID,
		Amount: req.Amount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        resp.Message,
		"updatedBalance": resp.UpdatedBalance,
	})
}

func (h *AccountHandler) ProcessTransaction(c *gin.Context) {
	var request struct {
		AccountID uint64  `json:"account_id"`
		Type      string  `json:"type"`  
		Amount    float64 `json:"amount"` 
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	transactionType, err := mapTransactionType(request.Type)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.ProcessTransaction(context.Background(), &pb.ProcessTransactionRequest{
		AccountId: request.AccountID,
		Type:      transactionType,
		Amount:    request.Amount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        resp.Message,
		"transaction_id": resp.TransactionId,
	})
}

func mapTransactionType(transactionType string) (pb.TransactionType, error) {
	switch transactionType {
	case "deposit":
		return pb.TransactionType_DEPOSIT, nil
	case "withdrawal":
		return pb.TransactionType_WITHDRAWAL, nil
	default:
		return 0, fmt.Errorf("invalid transaction type: %s", transactionType)
	}
}

func (h *AccountHandler) GetTransaction(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	resp, err := h.Client.GetTransaction(context.Background(), &pb.GetTransactionRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction": resp.Transaction})
}

func (h *AccountHandler) GetTransactionsByAccount(c *gin.Context) {
	accountIDStr := c.Param("account_id")
	accountID, err := strconv.ParseUint(accountIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	resp, err := h.Client.GetTransactionsByAccount(context.Background(), &pb.GetTransactionsByAccountRequest{AccountId: accountID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": resp.Transactions})
}