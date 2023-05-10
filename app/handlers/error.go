package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)
func Error() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
	 fmt.Println(err.Error())
	 response := fiber.Map{
		"message": err.Error(),
	 }
	 return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
 }
