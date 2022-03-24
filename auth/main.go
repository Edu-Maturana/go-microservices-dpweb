package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	microservice := fiber.New()

	microservice.Get("/api/auth", func(c *fiber.Ctx) error {
		return c.JSON("Hello from auth!")
	})

	fmt.Println("Server started on port 8091")
	microservice.Listen(":8091")
}
