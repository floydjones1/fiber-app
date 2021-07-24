package internal

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/floydjones1/fiber-app/internal/books"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"xorm.io/xorm"
)

func Start() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
		ReadTimeout:   time.Second * 5,
	})

	_, err := xorm.NewEngine("postgres", "localhost:5432")
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	// only fork under production mode
	// if !fiber.IsChild() {
	// 	fmt.Println("I'm the parent pcess")
	// } else {
	// 	fmt.Println("I'm a child process")
	// }

	app.Use(limiter.New(limiter.Config{
		Expiration: 30 * time.Second,
		Max:        3,
	}))
	app.Use(logger.New(logger.Config{
		Format:       "${yellow}[DEBUG] ${reset}- ${white}${time} ${reset}- ${magenta}${status} ${reset}- ${green}${method} ${reset}- ${white}${latency} ${reset}- ${cyan}${path}\n",
		TimeInterval: 500 * time.Millisecond,
		TimeZone:     "America/New_York",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	books.CreateService(app)

	if err := run(app); err != nil {
		fmt.Println(err)
	}
}

func run(app *fiber.App) error {
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP)

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	err := app.Shutdown()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
