package interfaces

import "github.com/ShahabazSultha/Trasactionlog/pkg/model"

type INotifRepo interface {
	CreateNewNotification(msg *model.TransactionLog) error
	GetNotificationsForUser(AccountId string, limit, offset int) ([]model.TransactionLog, error)
}
