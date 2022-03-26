package components

import (
	"ecommerce-ms/products/entitie"
	"ecommerce-ms/repository"
	"ecommerce-ms/utils"
)

var database = utils.GetEnvVar("DATABASE")
var mongoDbRepository, _ = repository.Connect()
var collection = mongoDbRepository.Database(database).Collection("products")

func GetAllProductsService() ([]interface{}, error) {
	results, err := repository.GetAll(collection)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetOneProductService(filter interface{}) (interface{}, error) {
	result, err := repository.GetOne(collection, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateProductService(product *entitie.Product) error {
	err := repository.InsertOne(collection, product)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProductService(filter interface{}, update interface{}) error {
	err := repository.UpdateOne(collection, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProductService(filter interface{}) error {
	err := repository.DeleteOne(collection, filter)
	if err != nil {
		return err
	}

	return nil
}
