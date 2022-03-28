package components

import (
	"ecommerce-ms/products/entitie"
	"ecommerce-ms/store"
	"ecommerce-ms/utils"

	"github.com/rs/xid"
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

func GetOneProductService(id string) (interface{}, error) {
	result, err := store.GetOne(collection, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateProductService(product *entitie.Product) error {
	product.ID = xid.New().String()
	err := store.InsertOne(collection, product)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProductService(id string, update interface{}) error {
	err := store.UpdateOne(collection, id, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProductService(id string) error {
	err := store.DeleteOne(collection, id)
	if err != nil {
		return err
	}

	return nil
}
