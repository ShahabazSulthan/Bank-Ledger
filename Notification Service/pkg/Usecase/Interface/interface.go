package interface_usecase

import "github.com/ShahabazSultha/Trasactionlog/pkg/model"

type INotifUseCase interface {
	KafkaMessageConsumer()
	GetNotificationsForUser(AccountId string, limit, offset int) ([]model.TransactionLog, error)
}
