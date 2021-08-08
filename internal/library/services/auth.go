package services

import (
	"fmt"
	"time"

	"github.com/floydjones1/fiber-app/internal/data"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	secret string = "secret"
)

type AuthService struct {
	Store data.Stores
}

type SignUpRequest struct {
	Email    string
	Password string
	Name     string
}

func (u *AuthService) SignUp(c *fiber.Ctx) error {
	req := new(SignUpRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return fiber.NewError(401, "please enter valid credentials")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(500, "please enter valid credentials")
	}
	user := &data.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hash),
	}

	if err := u.Store.User.InsertUser(user); err != nil {
		fmt.Println(err)
		return fiber.NewError(500, "failed to insert user")
	}

	return c.JSON(user)
}

type LoginRequest struct {
	Email    string
	Password string
}

func (u *AuthService) Login(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	if req.Email == "" || req.Password == "" {
		return fiber.NewError(401, "please enter credentials")
	}

	user, found, err := u.Store.User.GetUserByEmail(req.Email)
	if err != nil {
		return fiber.NewError(500, "failed to find user")
	}
	if !found {
		return fiber.NewError(404, "failed to verify credentials")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Minute * 30)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func (u *AuthService) Logout(c *fiber.Ctx) error {
	return nil
}
