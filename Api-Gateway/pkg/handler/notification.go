package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ShahabazSulthan/Api-Gateway/pkg/model"
	"github.com/ShahabazSulthan/Api-Gateway/pkg/pb"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	Client pb.NotificationServiceClient
}

func NewNotificationHandler(client pb.NotificationServiceClient) *NotificationHandler {
	return &NotificationHandler{Client: client}
}

func (h *NotificationHandler) GetNotificationsForUser(c *gin.Context) {
	accountID := c.Param("account_id")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset value"})
		return
	}

	req := &pb.RequestGetNotifications{
		AccountId: accountID,
		Limit:     limit,
		OffSet:    offset,
	}

	resp, err := h.Client.GetNotificationsForUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resp.ErrorMessage != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resp.ErrorMessage})
		return
	}

	notifications := make([]model.TransactionLog, len(resp.Notifications))
	for i, notif := range resp.Notifications {
		notifications[i] = model.TransactionLog{
			TransactionID:   notif.TransactionId,
			AccountID:       notif.AccountId,
			Amount:          notif.Amount,
			TransactionType: notif.TransactionType,
			Status:          notif.Status,
			CreatedAt:       notif.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}
