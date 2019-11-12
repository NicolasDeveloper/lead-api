package database

import (
	"context"
	"fmt"
	"log"

	"github.com/NicolasDeveloper/lead-api/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection : Providers the Mongo's connection
func Connection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://" + env.Parameters.Database.Ip + ":" + env.Parameters.Database.Port)
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
