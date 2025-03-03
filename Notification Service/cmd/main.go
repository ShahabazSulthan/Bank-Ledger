package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ShahabazSultha/Trasactionlog/pkg/config"
	"github.com/ShahabazSultha/Trasactionlog/pkg/di"
	"github.com/ShahabazSultha/Trasactionlog/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	notifServer, err := di.InitializeNotificationServer(config)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", config.PortMngr.RunnerPort)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Notification Service started on:", config.PortMngr.RunnerPort)
	defer lis.Close()

	grpcServer := grpc.NewServer()

	pb.RegisterNotificationServiceServer(grpcServer, notifServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start Notification_service server:%v", err)
	}
}
