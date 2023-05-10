package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/KayoRonald/golang-api/app/database"
	"github.com/KayoRonald/golang-api/app/handlers"
	"github.com/KayoRonald/golang-api/app/middleware"
	"github.com/KayoRonald/golang-api/app/router"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", router.GetUser)
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.Error(),
	})
	database.ConnectDB()
	// middleware cors
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
