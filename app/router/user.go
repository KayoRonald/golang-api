package router

import (
	"fmt"

	"github.com/KayoRonald/golang-api/app/database"
	"github.com/KayoRonald/golang-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// type Book struct {
// 	BookID      string `json:"book_id"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// }

// func CreateResponseUser(book model.Book) Book {
// 	return Book{BookID: book.BookID, Title: book.Title,Description: book.Description }
// }

// Handler
func GetUser(c *fiber.Ctx) error {
	// teste := ConnectDB.
	book := []model.Book{}
	result := database.Database.Db.First(&book)
	fmt.Print(result.RowsAffected)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": "Ops",
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
