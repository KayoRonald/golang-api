package router

import (
	"github.com/KayoRonald/golang-api/app/database"
	"github.com/KayoRonald/golang-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Handler
func GetBook(c *fiber.Ctx) error {
	book := []model.Book{}
	result := database.Database.Db.Find(&book)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Nenhum livro foi cadastrado",
			"status": "err",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": book,
		"status": "sucess",
	})
}

func ByIDGet(c *fiber.Ctx) error {
	id := c.Params("id")
	book := model.Book{}
	result := database.Database.Db.First(&book, "id = ?", id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Nenhum encontrado neste ID",
			"status": "err",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": book,
		"status": "sucess",
	})
}

func PostBook(c *fiber.Ctx) error {
	book := new(model.Book)
	book.BookID = uuid.New().String()
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status": "err",
		})
	}
	database.Database.Db.Create(book)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": book,
		"status": "sucess",
	})
}
