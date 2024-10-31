package middlewares

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RequestTimingMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)

	fmt.Printf(
		"%s: %vms\n",
		c.Path(),
		duration.Milliseconds(),
	)
	return err
}
