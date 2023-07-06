package controllers

import (
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
	"github.com/KayoRonald/golang-api/app/database"
	"github.com/KayoRonald/golang-api/app/models"
)

// Handler
func GetBook(c *fiber.Ctx) error {
	book := []models.Book{}
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
	book := models.Book{}
	result := database.Database.Db.First(&book, "id = ?", id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
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
	book := new(models.Book)
	book.ID = uuid.New().String()
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

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	book := models.Book{}
	result := database.Database.Db.Where("id = ?", id).Delete(&book)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Nenhum encontrado neste ID",
			"status": "err",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Livro deletado com sucesso!",
		"status": "sucess",
	})
}

func PutById(c *fiber.Ctx) error {
	id := c.Params("id")
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status": "err",
		})
	}
	result := database.Database.Db.Where("id = ?", id).Updates(&book)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Nenhum encontrado neste ID",
			"status": "err",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": book,
		"status": "sucess",
	})
}