package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection : Providers the Mongo's connection
func Connection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://165.22.15.231:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
