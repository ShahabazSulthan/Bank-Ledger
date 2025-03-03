package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ShahabazSultha/Trasactionlog/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDbCollection struct {
	TransactionLog *mongo.Collection
}

func ConnectDatabaseMongo(config *config.MongoDataBase) (*MongoDbCollection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	fmt.Println("----------connection uri:", config.MongoDbURL)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDbURL).SetServerAPIOptions(serverAPI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("can't ping to db, err:", err)
		return nil, fmt.Errorf("ping to MongoDB failed: %w", err)
	}

	fmt.Printf("\nconnected to MongoDB, on database %s\n", config.DataBase)

	var mongoCollections MongoDbCollection
	db := client.Database(config.DataBase)
	mongoCollections.TransactionLog = db.Collection("TransactionLog")

	return &mongoCollections, nil
}
