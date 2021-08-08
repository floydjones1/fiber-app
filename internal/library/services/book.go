package services

import (
	"fmt"

	"github.com/floydjones1/fiber-app/internal/data"
	"github.com/gofiber/fiber/v2"
)

type BookService struct {
	Store data.Stores
}

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
