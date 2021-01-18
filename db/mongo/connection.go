package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitClient(uri string) (err error) {
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return
	}

	log.Println("Connected to MongoDB!")
	return
}

func GetClient() *mongo.Client {
	return client
}

func GetDatabase() *mongo.Database {
	return client.Database("test")
}

func GetCollection() *mongo.Collection {
	return GetDatabase().Collection("User")
}

func Close() error {
	if client == nil {
		return fmt.Errorf("Client already closed")
	}
	err := client.Disconnect(context.TODO())
	return err
}
