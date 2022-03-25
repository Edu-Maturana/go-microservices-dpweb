package main

import (
	"ecommerce-ms/store"
	"testing"
)

func TestGetEnvVar(t *testing.T) {
	_, err := store.Connect()

	if err != nil {
		t.Errorf("Error connecting to MongoDB: %s", err)
	}
}
