package server

import (
	"context"
	"fmt"

	interface_usecase "github.com/ShahabazSultha/Trasactionlog/pkg/Usecase/Interface"
	"github.com/ShahabazSultha/Trasactionlog/pkg/pb"
)

type NotifService struct {
	NotifUseCase interface_usecase.INotifUseCase
	pb.NotificationServiceServer
}

func NewNotifServiceServer(notifUseCase interface_usecase.INotifUseCase) *NotifService {
	return &NotifService{NotifUseCase: notifUseCase}
}

func (u *NotifService) GetNotificationsForUser(ctx context.Context, req *pb.RequestGetNotifications) (*pb.ResponseGetNotifications, error) {
	notificationsData, err := u.NotifUseCase.GetNotificationsForUser(req.AccountId, int(req.Limit), int(req.OffSet))
	if err != nil {
		return &pb.ResponseGetNotifications{
			ErrorMessage: err.Error(),
		}, nil
	}

	var notifications []*pb.TransactionLog

	for _, data := range notificationsData {
		notifications = append(notifications, &pb.TransactionLog{
			TransactionId:   data.TransactionID,
			AccountId:       data.AccountID,
			Amount:          data.Amount,
			TransactionType: data.TransactionType,
			Status:          data.Status,
			CreatedAt:       data.CreatedAt.String(),
		})
	}

	fmt.Println("Notif", notifications)

	return &pb.ResponseGetNotifications{
		Notifications: notifications,
	}, nil
}
