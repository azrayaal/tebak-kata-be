package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define the global database name variable
var databaseName = "teka_apps"

// func ConnectToMongoDB(databaseName string) (*mongo.Client, *mongo.Database, error) {
func ConnectToMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://azrayaal:rahasia@teka.cf0mmqo.mongodb.net/?retryWrites=true&w=majority&appName=teka")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to ensure connectivity
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}
