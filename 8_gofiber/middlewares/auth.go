package middlewares

import (
	"gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	// armazena um valor no Locals
	c.Locals("request_date", utils.GetDateNowFormated())

	if token != "valid-token" {
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{"error": "User unauthorized!"})
	}

	return c.Next()
}
