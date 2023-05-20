package router

import (
	// "fmt"

	"github.com/KayoRonald/golang-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	// "gorm.io/gorm"
)

// func CreateUsers(db *gorm.DB, Title string, UrlImage string, Description string, Price int) error {
// 	return db.Create(&model.Book{BookID: uuid.New().String(), Title: Title, Description: Description, UrlImage: UrlImage, Price: uint(Price)}).Error
// }

// Handler
func GetUser(c *fiber.Ctx) error {
	// teste := ConnectDB.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello World",
	})
}

func PostUser(c *fiber.Ctx) error {
	// getting user if no error
	book := new(model.Book)
	book.BookID = uuid.New().String()
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": book,
	})
}
