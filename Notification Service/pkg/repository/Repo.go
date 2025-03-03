package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ShahabazSultha/Trasactionlog/pkg/model"
	"github.com/ShahabazSultha/Trasactionlog/pkg/repository/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NotifRepo struct {
	Collection *mongo.Collection
}

func NewNotifRepo(collection *mongo.Collection) interfaces.INotifRepo {
	return &NotifRepo{Collection: collection}
}

func (n *NotifRepo) CreateNewNotification(msg *model.TransactionLog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := n.Collection.InsertOne(ctx, msg)
	return err
}

func (n *NotifRepo) GetNotificationsForUser(AccountID string, limit, offset int) ([]model.TransactionLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"accountid": AccountID}
	opts := options.Find().
		SetSort(bson.M{"created_at": -1}).
		SetLimit(int64(limit)).
		SetSkip(int64(offset))

	cursor, err := n.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notifications []model.TransactionLog
	if err = cursor.All(ctx, &notifications); err != nil {
		return nil, err
	}


	fmt.Println("Notificattttttttion", notifications)

	return notifications, nil
}
