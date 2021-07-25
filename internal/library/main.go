package library

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/floydjones1/fiber-app/config"
	"github.com/floydjones1/fiber-app/internal/data"
	"github.com/floydjones1/fiber-app/internal/library/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Start() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config/local.yml", "path to config file")
	flag.Parse()

	config, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Debug().Msg("Loaded config....Starting up server")

	app := createApp(config.Server)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	db, err := data.InitializeDB(config.Database)
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

func createApp(serverConf config.Server) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:      serverConf.Prefork,
		ServerHeader: "Test App v1.0.0",
		AppName:      "Fiber",
		ReadTimeout:  time.Second * time.Duration(serverConf.ReadTimeout),
	})
	// only fork under production mode
	// if !fiber.IsChild() {

	logOutput := "${yellow}[DEBUG] ${reset}- " +
		"${white}${time} ${reset}- ${magenta}${status} ${reset}- " +
		"${green}${method} ${reset}- ${white}${latency} ${reset}- ${cyan}${path}\n"
	app.Use(logger.New(logger.Config{
		Format:   logOutput,
		TimeZone: "America/New_York",
	}))

	return app
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
