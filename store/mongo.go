package store

import (
	"context"
	"ecommerce-ms/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var documents []interface{}
	for cursor.Next(context.TODO()) {
		var document bson.M
		err := cursor.Decode(&document)
		if err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}

	return documents, nil
}

func GetOne(collection *mongo.Collection, id string) (interface{}, error) {
	result := collection.FindOne(context.TODO(), bson.M{"_id": id})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var document bson.M
	err := result.Decode(&document)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func InsertOne(collection *mongo.Collection, document interface{}) error {
	_, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOne(collection *mongo.Collection, id string, update interface{}) error {
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOne(collection *mongo.Collection, id string) error {
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
