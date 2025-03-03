package interface_kafka

import "github.com/ShahabazSulthan/BankingService/pkg/domain"

type IKafkaProducer interface {
	KafkaNotificationProducer(message *domain.TransactionLog) error
}


