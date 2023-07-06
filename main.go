package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/KayoRonald/golang-api/app/configs"
	"github.com/KayoRonald/golang-api/app/database"
	"github.com/KayoRonald/golang-api/app/handlers"
	"github.com/KayoRonald/golang-api/app/middleware"
	"github.com/KayoRonald/golang-api/app/router"
)

var (
	logs *configs.Logger
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", router.GetBook)
	app.Get("/:id", router.ByIDGet)
	app.Post("/", router.PostBook)
}

func main() {
	// inialize database
	database.ConnectDB()
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	logs = configs.GetLogger("run")
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.Error(),
	})
	app.Use(logger.New())
	app.Use(middleware.CorsMiddleware())
	app.Use(middleware.Limiter())
	setUpRoutes(app)
	app.Use(handlers.NotFound)

	port := os.Getenv("port")
	if port == "" {
		logs.Info("Listen on environment specified PORT, '3000' otherwise")
		port = "3000"
	}
	app.Listen(":" + port)
	// log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
	return nil
}
