package usecase

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	interface_usecase "github.com/ShahabazSultha/Trasactionlog/pkg/Usecase/Interface"
	"github.com/ShahabazSultha/Trasactionlog/pkg/config"
	"github.com/ShahabazSultha/Trasactionlog/pkg/model"
	"github.com/ShahabazSultha/Trasactionlog/pkg/repository/interfaces"
)

type NotifUseCase struct {
	Repo        interfaces.INotifRepo
	KafkaConfig config.KafkaConfigs
}

func NewNotifUseCase(notifRepo interfaces.INotifRepo, config config.KafkaConfigs) interface_usecase.INotifUseCase {
	return &NotifUseCase{
		Repo:        notifRepo,
		KafkaConfig: config,
	}
}

func (n *NotifUseCase) KafkaMessageConsumer() {
	fmt.Println("--------- Kafka consumer initiated ---------")

	// Configure Sarama settings
	configs := sarama.NewConfig()
	configs.Consumer.Return.Errors = true
	configs.Version = sarama.V2_1_0_0

	// Initialize Kafka consumer
	consumer, err := sarama.NewConsumer([]string{n.KafkaConfig.KafkaPort}, configs)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
		return
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Printf("Failed to close Kafka consumer: %v", err)
		}
	}()

	// Set up partition consumer
	partitionConsumer, err := consumer.ConsumePartition(n.KafkaConfig.KafkaTopicNotification, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to consume partition: %v", err)
		return
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Printf("Failed to close partition consumer: %v", err)
		}
	}()

	// Listen for messages
	for message := range partitionConsumer.Messages() {
		if message == nil {
			log.Println("Received nil message from Kafka, skipping...")
			continue
		}

		var msg model.TransactionLog
		if err := json.Unmarshal(message.Value, &msg); err != nil {
			log.Printf("Failed to unmarshal Kafka message: %v", err)
			continue
		}
		fmt.Printf("Received message: %+v\n", msg)

		// Store notification in the repository
		if err := n.Repo.CreateNewNotification(&msg); err != nil {
			log.Printf("Failed to create notification: %v", err)
		}
	}
}

func (n *NotifUseCase) GetNotificationsForUser(accountId string, limit, offset int) ([]model.TransactionLog, error) {
	// Retrieve notifications from the repository
	notifData, err := n.Repo.GetNotificationsForUser(accountId, limit, offset)
	if err != nil {
		return nil, err
	}

	return notifData, nil
}
