package components

import (
	"ecommerce-ms/products/entitie"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

func GetProductsController(c *fiber.Ctx) error {
	products, err := GetAllProductsService()
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.JSON(products)
}

func GetProductController(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := GetOneProductService(id)
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.JSON(product)
}

func CreateProductController(c *fiber.Ctx) error {
	var newProduct entitie.Product
	if err := c.BodyParser(&newProduct); err != nil {
		return c.Status(400).JSON(err)
	}

	if err := validator.New().Struct(newProduct); err != nil {
		return c.Status(400).JSON(err)
	}

	if err := CreateProductService(&newProduct); err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(201).JSON("Product created successfully")
}

func UpdateProductController(c *fiber.Ctx) error {
	id := c.Params("id")
	var product entitie.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err)
	}

	if err := validator.New().Struct(product); err != nil {
		return c.Status(400).JSON(err)
	}

	if err := UpdateProductService(id, &product); err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON("Product updated successfully")
}

func DeleteProductController(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := DeleteProductService(id); err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON("Product deleted successfully")
}
