package client

import (
	"fmt"

	"github.com/ShahabazSulthan/Api-Gateway/pkg/config"
	"github.com/ShahabazSulthan/Api-Gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAccountClient(Config *config.Config) (*pb.AccountServiceClient, error) {
	cc, err := grpc.Dial(Config.BankSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error Connect auth Clinet : ", err)
		return nil, err
	}

	client := pb.NewAccountServiceClient(cc)

	return &client, nil
}
