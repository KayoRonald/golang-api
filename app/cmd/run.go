package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/KayoRonald/golang-api/app/configs"
	"github.com/KayoRonald/golang-api/app/handlers"
	"github.com/KayoRonald/golang-api/app/middleware"
	"github.com/KayoRonald/golang-api/app/router"
)

var (
	logs *configs.Logger
)

func Run() error {
	logs = configs.GetLogger("run")
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.Error(),
	})
	app.Use(logger.New())
	app.Use(middleware.CorsMiddleware())
	app.Use(middleware.Limiter())
	router.BookRoutes(app)
	app.Use(handlers.NotFound)

	port := os.Getenv("port")
	if port == "" {
		logs.Info("Environment specified PORT, '3000' otherwise")
		port = "3000"
	}
	app.Listen(":" + port)
	return nil
}
