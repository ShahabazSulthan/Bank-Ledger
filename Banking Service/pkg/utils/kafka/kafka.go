package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/ShahabazSulthan/BankingService/pkg/config"
	"github.com/ShahabazSulthan/BankingService/pkg/domain"
	interface_kafka "github.com/ShahabazSulthan/BankingService/pkg/utils/kafka/interfaces"
)

type KafkaProducer struct {
	Config config.KafkaConfigs
}

func NewKafkaProducer(config config.KafkaConfigs) interface_kafka.IKafkaProducer {
	return &KafkaProducer{Config: config}
}

func CheckKafkaConnection(brokers []string, configs *sarama.Config) error {
	log.Println("Checking Kafka connection...")
	client, err := sarama.NewClient(brokers, configs)
	if err != nil {
		log.Println("Kafka connection failed:", err)
		return err
	}
	defer client.Close()
	log.Println("Kafka connection successful.")
	return nil
}

func (k *KafkaProducer) KafkaNotificationProducer(message *domain.TransactionLog) error {
	fmt.Println("Sending message to Kafka:", message)

	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{k.Config.KafkaPort}, configs)
	if err != nil {
		log.Printf("Error creating Kafka producer: %v\n", err)
		return err
	}
	defer producer.Close()

	msgJson, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling message to JSON: %v\n", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: k.Config.KafkaTopicNotification,
		Key:   sarama.StringEncoder(message.AccountID),
		Value: sarama.ByteEncoder(msgJson),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Error sending message to Kafka topic '%s': %v\n", k.Config.KafkaTopicNotification, err)
		return err
	}

	log.Printf("[Kafka Producer] Message sent successfully | Partition: %d | Offset: %d\n", partition, offset)
	return nil
}
