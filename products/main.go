package main

import (
	"ecommerce-ms/products/components"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	microservice := fiber.New()

	microservice.Get("/api/products", components.GetProductController)
	microservice.Get("/api/products/:id", components.GetProductController)
	microservice.Post("/api/products", components.CreateProductController)
	microservice.Put("/api/products/:id", components.UpdateProductController)
	microservice.Delete("/api/products/:id", components.DeleteProductController)

	fmt.Println("Server started on port 8080")
	microservice.Listen(":8080")
}
