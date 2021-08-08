// This is custom goose binary with sqlite3 support only.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/floydjones1/fiber-app/config"
	_ "github.com/floydjones1/fiber-app/internal/data/migrations"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

var (
	flags      = flag.NewFlagSet("goose", flag.ExitOnError)
	dir        = flags.String("dir", ".", "directory with migration files")
	configPath = flag.String("config", "./config/local.yml", "path to config file")
)

// ./goose -config=./dir/to/db/config -dir=./dir/to/migrations
func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	config, err := config.GetConfig(*configPath)
	if err != nil {
		log.Fatalf("goose: failed to read DB config: %v\n", err)
		return
	}

	dbConfig := config.Database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.DatabaseName)

	db, err := goose.OpenDBWithDriver("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	command := args[0]
	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
