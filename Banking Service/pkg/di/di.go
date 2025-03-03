package di

import (
	"fmt"
	"log"

	server "github.com/ShahabazSulthan/BankingService/pkg/Server"
	"github.com/ShahabazSulthan/BankingService/pkg/config"
	"github.com/ShahabazSulthan/BankingService/pkg/db"
	"github.com/ShahabazSulthan/BankingService/pkg/repository"
	"github.com/ShahabazSulthan/BankingService/pkg/usecase"
	hash "github.com/ShahabazSulthan/BankingService/pkg/utils/hashed_password"
	"github.com/ShahabazSulthan/BankingService/pkg/utils/kafka"
)

func InitializeAuthService(config *config.Config) (*server.AccountServer, error) {
	hashUtil := hash.NewHashUtil()
	DB, err := db.ConnectDatabase(&config.DB, hashUtil)
	if err != nil {
		fmt.Println("Error in connecting database from Dependency Injection")
		return nil, err
	}

	kafkaProducer := kafka.NewKafkaProducer(config.Kafka)
	if err := kafka.CheckKafkaConnection([]string{config.Kafka.KafkaPort}, nil); err != nil {
		log.Println("Failed to connect to Kafka:", err)
		return nil, err
	}
	fmt.Println("Kafka connection established successfully.")

	AccountRepo := repository.NewAccountRepo(DB)
	AccountUseCase := usecase.NewAccountUseCase(AccountRepo, kafkaProducer)
	AccountServer := server.NewAccountServer(AccountUseCase)

	return AccountServer, nil
}
