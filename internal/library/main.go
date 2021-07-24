package library

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/floydjones1/fiber-app/internal/data"
	"github.com/floydjones1/fiber-app/internal/library/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Start() {
	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Debug().Msg("This message appears only when log level set to Debug")
	log.Info().Msg("This message appears when log level set to Debug or Info")

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Test App v1.0.1",
		// ErrorHandler: ,
		AppName:     "Fiber",
		ReadTimeout: time.Second * 5,
	})
	// only fork under production mode
	// if !fiber.IsChild() {
	// 	fmt.Println("I'm the parent process")
	// } else {
	// 	fmt.Println("I'm a child process")
	// }
	// app.Use(limiter.New(limiter.Config{
	// 	Expiration: 30 * time.Second,
	// 	Max:        3,
	// }))

	logOutput := "${yellow}[DEBUG] ${reset}- " +
		"${white}${time} ${reset}- ${magenta}${status} ${reset}- " +
		"${green}${method} ${reset}- ${white}${latency} ${reset}- ${cyan}${path}\n"
	app.Use(logger.New(logger.Config{
		Format:       logOutput,
		TimeInterval: 500 * time.Millisecond,
		TimeZone:     "America/New_York",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	db, err := data.InitializeDB()
	if err != nil {
		log.Err(err)
		return
	}

	routes.SetupHandlers(app, routes.Options{
		Store: db,
	})

	if err := run(app); err != nil {
		log.Err(err)
		return
	}
}

func run(app *fiber.App) error {
	go func() {
		if err := app.Listen(":8000"); err != nil {
			log.Err(err)
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
