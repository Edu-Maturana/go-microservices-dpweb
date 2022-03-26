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

type MongoRepository interface {
	GetAll(collection *mongo.Collection) ([]interface{}, error)
	GetOne(collection *mongo.Collection, filter interface{}) (interface{}, error)
	InsertOne(collection *mongo.Collection, document interface{}) error
	UpdateOne(collection *mongo.Collection, filter interface{}, update interface{}) error
	DeleteOne(collection *mongo.Collection, filter interface{}) error
}

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
	var results []interface{}

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func GetOne(collection *mongo.Collection, filter interface{}) (interface{}, error) {
	result := collection.FindOne(context.TODO(), filter)
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
