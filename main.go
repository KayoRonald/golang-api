package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KayoRonald/golang-api/app/handlers"
	"github.com/KayoRonald/golang-api/app/middleware"
	"github.com/KayoRonald/golang-api/app/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.Error(),
	})
	// middleware cors
	app.Use(middleware.CorsMiddleware())
	// middleware limiter request
	app.Use(middleware.Limiter())

	app.Get("/", router.GetUser)
	// middleware not found
	app.Use(handlers.NotFound)

	// Listen on environment specified PORT, "3000" otherwise
	port := os.Getenv("port")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
