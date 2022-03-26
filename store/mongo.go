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

func GetAll(collection *mongo.Collection) ([]interface{}, error) {
	cursor, err := collection.Find(context.Background(), nil, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []interface{}
	for cursor.Next(context.Background()) {
		var result interface{}
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func GetOne(collection *mongo.Collection, filter interface{}) (interface{}, error) {
	result := collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var res interface{}
	err := result.Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func InsertOne(collection *mongo.Collection, document interface{}) error {
	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOne(collection *mongo.Collection, filter interface{}, update interface{}) error {
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOne(collection *mongo.Collection, filter interface{}) error {
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
