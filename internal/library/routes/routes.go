package routes

import (
	"fmt"

	"github.com/floydjones1/fiber-app/internal/data"
	"github.com/gofiber/fiber/v2"
)

type BookService struct {
	Store data.Stores
}
type Options struct {
	Store *data.Stores
}

func SetupHandlers(app *fiber.App, opt Options) {
	bookSvc := BookService{
		Store: *opt.Store,
	}
	group := app.Group("/books")
	group.Get("/", bookSvc.GetBooks)
	group.Get("/error", bookSvc.Error)

}

func (b *BookService) GetBooks(c *fiber.Ctx) error {
	fmt.Println("here we go")
	b.Store.UserStore.GetUser(2)
	return c.Status(200).JSON(&fiber.Map{
		"success": false,
		"error":   "There are no posts!",
	})

}

func (b *BookService) Error(c *fiber.Ctx) error {

	return fiber.NewError(fiber.StatusServiceUnavailable, "Books on vacation!")
}
