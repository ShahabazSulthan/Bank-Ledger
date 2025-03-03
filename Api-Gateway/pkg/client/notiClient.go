package client

import (
	"fmt"

	"github.com/ShahabazSulthan/Api-Gateway/pkg/config"
	"github.com/ShahabazSulthan/Api-Gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitNotificationClient(Config *config.Config) (*pb.NotificationServiceClient, error) {
	cc, err := grpc.Dial(Config.NotifsvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error Connect auth Clinet : ", err)
		return nil, err
	}

	client := pb.NewNotificationServiceClient(cc)

	return &client, nil
}
