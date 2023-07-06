package router

import (
	"github.com/KayoRonald/golang-api/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	app.Get("/", controllers.GetBook)
	app.Get("/:id", controllers.ByIDGet)
	app.Post("/", controllers.PostBook)
	app.Put("/:id", controllers.PutById)
	app.Delete("/:id", controllers.Delete)
}