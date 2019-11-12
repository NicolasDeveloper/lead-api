package leads

import (
	"context"
	"fmt"
	"log"

	"github.com/NicolasDeveloper/lead-api/database"
	"github.com/NicolasDeveloper/lead-api/dtos"
)

func Save(item dtos.Lead) error {
	client := database.Connection()
	collection := client.Database("easymile").Collection("leads")
	insertResult, err := collection.InsertOne(context.TODO(), item)

	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")

	if err != nil {
		return err
	}

	return nil
}
