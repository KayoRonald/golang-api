package router

import (
	"github.com/gofiber/fiber/v2"
)

// Handler
func GetUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Hello World",
	})
}
