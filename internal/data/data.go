package data

import (
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type Stores struct {
	DB *xorm.Engine
	UserStore
	BookStore
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbName   = "library"
)

func InitializeDB() (*Stores, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	engine, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL()

	if err := engine.Ping(); err != nil {
		return nil, err
	}

	stores := &Stores{
		DB:        engine,
		UserStore: UserStore{db: engine},
	}

	return stores, nil
}
