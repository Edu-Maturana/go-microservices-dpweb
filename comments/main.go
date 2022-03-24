package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	microservice := fiber.New()

	microservice.Get("/api/comments", func(c *fiber.Ctx) error {
		return c.JSON("Hello from comments!")
	})

	fmt.Println("Server started on port 8081")
	microservice.Listen(":8081")
}
