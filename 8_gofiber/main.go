package main

import (
	"fmt"
	"gofiber/handlers"
	"gofiber/middlewares"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Application is healthy: %s", time.Now()),
		})
	})

	// usando middlewares
	app.Use(middlewares.AuthMiddleware)
	app.Use(middlewares.ConfigureCors())

	app.Get("/api/users/user", handlers.GetUserByLocals)
	app.Get("/api/users/:username", handlers.GetUsernameHandler)
	app.Post("/api/products", middlewares.RequestTimingMiddleware, handlers.CreateProductHandler) // usando middleware em uma rota específica

	app.Listen(":8080")
}
