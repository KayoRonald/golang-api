package router

import (
	"github.com/KayoRonald/golang-api/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	v1 := app.Group("book")
	v1.Get("/", controllers.GetBook)
	v1.Get("/:id", controllers.ByIDGet)
	v1.Post("/", controllers.PostBook)
	v1.Put("/:id", controllers.PutById)
	v1.Delete("/:id", controllers.Delete)
}