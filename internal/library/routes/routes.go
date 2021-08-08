package routes

import (
	"github.com/floydjones1/fiber-app/internal/data"
	"github.com/floydjones1/fiber-app/internal/library/services"
	"github.com/gofiber/fiber/v2"
)

type Options struct {
	Store *data.Stores
}

func SetupHandlers(app *fiber.App, opt Options) {
	bookSvc := services.BookService{
		Store: *opt.Store,
	}

	authSvc := services.AuthService{
		Store: *opt.Store,
	}

	bookGroup := app.Group("/books")
	bookGroup.Get("/", bookSvc.GetBooks)
	bookGroup.Get("/error", bookSvc.Error)

	authGroup := app.Group("/auth")
	authGroup.Post("/signup", authSvc.SignUp)
	authGroup.Post("/login", authSvc.Login)
	authGroup.Post("/logout", authSvc.Logout)
}
