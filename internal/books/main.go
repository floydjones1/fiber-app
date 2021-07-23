package books

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type BookService struct{}

func (b *BookService) GetBooks(c *fiber.Ctx) error {
	fmt.Println("here we go")
	return c.Status(200).JSON(&fiber.Map{
		"success": false,
		"error":   "There are no posts!",
	})

}

func (b *BookService) Error(c *fiber.Ctx) error {

	return fiber.NewError(fiber.StatusServiceUnavailable, "Books on vacation!")
}

func CreateService(app *fiber.App) {
	bookSvc := BookService{}
	group := app.Group("/books")
	group.Get("/", bookSvc.GetBooks)
	group.Get("/error", bookSvc.Error)

}
