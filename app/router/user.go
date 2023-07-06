package router

import (

	"github.com/KayoRonald/golang-api/app/database"
	"github.com/KayoRonald/golang-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Handler
func GetUser(c *fiber.Ctx) error {
	book := []model.Book{}
	result := database.Database.Db.Find(&book)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Nenhum livro foi cadastrado",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": book,
	})
}

func PostUser(c *fiber.Ctx) error {
	book := new(model.Book)
	book.BookID = uuid.New().String()
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "err",
		})
	}
	database.Database.Db.Create(book)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": book,
	})
}
