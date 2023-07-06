package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/KayoRonald/golang-api/app/database"
	"github.com/KayoRonald/golang-api/app/handlers"
	"github.com/KayoRonald/golang-api/app/middleware"
	"github.com/KayoRonald/golang-api/app/router"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", router.GetUser)
	app.Post("/", router.PostUser)
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.Error(),
	})
	// inialize database
	database.ConnectDB()
	// middleware cors
	app.Use(logger.New())
	app.Use(middleware.CorsMiddleware())
	// middleware limiter request
	app.Use(middleware.Limiter())
	setUpRoutes(app)
	// middleware not found
	app.Use(handlers.NotFound)

	// Listen on environment specified PORT, "3000" otherwise
	port := os.Getenv("port")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
