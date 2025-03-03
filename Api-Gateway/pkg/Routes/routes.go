package routes

import (
	"github.com/ShahabazSulthan/Api-Gateway/pkg/handler"
	"github.com/gin-gonic/gin"
)

func AccountHandlerRoutes(router *gin.Engine, accountHandler *handler.AccountHandler, notifHandler *handler.NotificationHandler) {
	accountRoutes := router.Group("/accounts")
	{
		accountRoutes.POST("", accountHandler.CreateAccount)
		accountRoutes.GET("/:id", accountHandler.GetAccount)
		accountRoutes.PATCH("/:id/update-balance", accountHandler.UpdateBalance)
	}

	transactionRoutes := router.Group("/transactions")
	{
		transactionRoutes.POST("", accountHandler.ProcessTransaction)
		transactionRoutes.GET("/:id", accountHandler.GetTransaction)
		transactionRoutes.GET("/account/:account_id", accountHandler.GetTransactionsByAccount)
	}

	notificationRoutes := router.Group("/notifications")
	{
		notificationRoutes.GET("/:account_id", notifHandler.GetNotificationsForUser)
	}
}
