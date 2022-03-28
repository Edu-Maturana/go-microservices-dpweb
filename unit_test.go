package main

import (
	"ecommerce-ms/products/components"
	"ecommerce-ms/products/entitie"
	"ecommerce-ms/store"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

// DATABASE TESTS
func TestDBConnection(t *testing.T) {
	_, err := store.Connect()

	if err != nil {
		t.Errorf("Error connecting to DB: %s", err)
	}
}

// PRODUCTS SERVICES TESTS
func TestGetAllProducts(t *testing.T) {
	results, err := components.GetAllProductsService()

	if err != nil {
		t.Errorf("Error getting all products: %s", err)
	}

	if len(results) == 0 {
		t.Error("No products found")
	}
}

func TestGetOneProduct(t *testing.T) {
	id := "c8vb8c9ikvt6r99s9u00"
	result, err := components.GetOneProductService(id)

	if err != nil {
		t.Errorf("Error getting one product: %s", err)
	}

	if result == nil {
		t.Error("Product not found")
	}
}

func TestCreateProduct(t *testing.T) {
	product := &entitie.Product{
		ID:          "1",
		Name:        "Test Product",
		Description: "Test Product Description",
		Image:       "http://test.com/image.png",
		Stock:       10,
		Price:       10,
	}

	err := components.CreateProductService(product)

	if err != nil {
		t.Errorf("Error creating product: %s", err)
	}
}

func TestUpdateProduct(t *testing.T) {
	id := "1"

	err := components.UpdateProductService(id, bson.M{"$set": bson.M{"name": "Test Product Updated"}})

	if err != nil {
		t.Errorf("Error updating product: %s", err)
	}
}

func TestDeleteProduct(t *testing.T) {
	id := "1"

	err := components.DeleteProductService(id)

	if err != nil {
		t.Errorf("Error deleting product: %s", err)
	}
}
