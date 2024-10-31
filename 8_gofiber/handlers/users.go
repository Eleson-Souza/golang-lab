package handlers

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GetUsernameHandler(c *fiber.Ctx) error {
	userName := c.Params("username")

	userNameDecoded, err := url.QueryUnescape(userName)

	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, "username param invalid: ", err.Error())
	}

	message := fmt.Sprintf("Hello, %s!", userNameDecoded)

	return c.JSON(
		fiber.Map{
			"message": message,
		},
	)
}

func GetUserByLocals(c *fiber.Ctx) error {
	// Recupera o valor do Locals
	requestDate := c.Locals("request_date")

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Request executed on %s", requestDate),
	})
}
