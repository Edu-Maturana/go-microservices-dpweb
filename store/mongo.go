package store

import (
	"context"
	"ecommerce-ms/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI = utils.GetEnvVar("MONGO_DB_ATLAS_URI")

func Connect() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB", err)
		return nil, err
	} else {
		log.Println("Connected to MongoDB")
	}

	return client, nil
}
