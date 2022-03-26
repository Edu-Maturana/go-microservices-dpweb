package main

import (
	"ecommerce-ms/store"
	"testing"
)

func TestDBConnection(t *testing.T) {
	_, err := store.Connect()

	if err != nil {
		t.Errorf("Error connecting to MongoDB: %s", err)
	}
}
