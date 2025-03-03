package di

import (
	"fmt"

	"github.com/ShahabazSultha/Trasactionlog/pkg/config"
	"github.com/ShahabazSultha/Trasactionlog/pkg/db"
	"github.com/ShahabazSultha/Trasactionlog/pkg/repository"
	"github.com/ShahabazSultha/Trasactionlog/pkg/server"
	usecase "github.com/ShahabazSultha/Trasactionlog/pkg/Usecase"
)

func InitializeNotificationServer(config *config.Config) (*server.NotifService, error) {
	dbInstance, err := db.ConnectDatabaseMongo(&config.MongoDB)
	if err != nil {
		fmt.Println("Error connecting to MongoDB in InitializeNotificationServer")
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	notifRepo := repository.NewNotifRepo(dbInstance.TransactionLog)
	notifUseCase := usecase.NewNotifUseCase(notifRepo, config.Kafka)

	go notifUseCase.KafkaMessageConsumer()

	return server.NewNotifServiceServer(notifUseCase), nil
}
