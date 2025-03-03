package di

import (
	"log"

	routes "github.com/ShahabazSulthan/Api-Gateway/pkg/Routes"
	"github.com/ShahabazSulthan/Api-Gateway/pkg/client"
	"github.com/ShahabazSulthan/Api-Gateway/pkg/config"
	"github.com/ShahabazSulthan/Api-Gateway/pkg/handler"
	"github.com/gin-gonic/gin"
)

func InitBankClient(router *gin.Engine, config *config.Config) error {
	Accountclient, err := client.InitAccountClient(config)
	if err != nil {
		log.Fatal(err)
	}

	Notficlient, err := client.InitNotificationClient(config)
	if err != nil {
		log.Fatal(err)
	}

	AccountHandler := handler.NewAccountHandler(*Accountclient)

	NotificationHandler := handler.NewNotificationHandler(*Notficlient)


	routes.AccountHandlerRoutes(router, AccountHandler,NotificationHandler)

	return nil
}
