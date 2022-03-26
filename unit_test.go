package main

import (
	"ecommerce-ms/repository"
	"testing"
)

func TestGetEnvVar(t *testing.T) {
	_, err := repository.Connect()

	if err != nil {
		t.Errorf("Error connecting to MongoDB: %s", err)
	}
}
