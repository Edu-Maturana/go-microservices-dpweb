package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	microservice := fiber.New()

	microservice.Get("/api/orders", func(c *fiber.Ctx) error {
		return c.JSON("Hello from orders!")
	})

	fmt.Println("Server started on port 8090")
	microservice.Listen(":8090")
}
