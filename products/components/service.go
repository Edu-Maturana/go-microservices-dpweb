package components

import (
	"ecommerce-ms/products/entitie"
	"ecommerce-ms/store"
	"ecommerce-ms/utils"
)

var database = utils.GetEnvVar("DATABASE")

var mongoDbStore, _ = store.Connect()
var collection = mongoDbStore.Database(database).Collection("products")

func GetAllProductsService() ([]interface{}, error) {
	results, err := store.GetAll(collection)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetOneProductService(filter interface{}) (interface{}, error) {
	result, err := store.GetOne(collection, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateProductService(product *entitie.Product) error {
	err := store.InsertOne(collection, product)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProductService(filter interface{}, update interface{}) error {
	err := store.UpdateOne(collection, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProductService(filter interface{}) error {
	err := store.DeleteOne(collection, filter)
	if err != nil {
		return err
	}

	return nil
}
