package services

import (
	"fmt"

	"github.com/floydjones1/fiber-app/internal/data"
	"github.com/gofiber/fiber/v2"
)

type AuthService struct {
	Store data.Stores
}

func (u *AuthService) SignUp(c *fiber.Ctx) error {
	return nil
}

type LoginRequest struct {
	Name     string
	Email    string
	Password string
}

type LoginResponse struct {
	Success bool
	Temp    map[string]int `json:"temp,omitempty"`
}

func (u *AuthService) Login(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	user, err := u.Store.User.GetUserByEmail(req.Email)
	if err != nil {
		return fiber.NewError(500, "had trouble finding user")
	}

	fmt.Println(user, *req)

	res := LoginResponse{Success: true, Temp: map[string]int{}}

	return c.JSON(res)
}

func (u *AuthService) Logout(c *fiber.Ctx) error {
	return nil
}
