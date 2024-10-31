package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateProductHandler(c *fiber.Ctx) error {
	//loading
	time.Sleep(time.Second * 5)

	c.Status(fiber.StatusCreated)

	return c.JSON(fiber.Map{
		"message": "Product registered with success",
	})
}
