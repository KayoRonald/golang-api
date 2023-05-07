package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// middleware cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(201).JSON(fiber.Map{
			"success": true,
			"message": "Hello World",
		})
	})

	// middleware not found
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "not found :/",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
