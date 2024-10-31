package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ConfigureCors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "https://example.com", // permitir apenas essa origem
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, OPTIONS",
	})
}
