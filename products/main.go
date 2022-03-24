package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	microservice := fiber.New()

	microservice.Get("/api/products", func(c *fiber.Ctx) error {
		return c.JSON("Hello from products!")
	})

	fmt.Println("Server started on port 8080")
	microservice.Listen(":8080")
}
