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
	pghost     = "localhost"
	pgport     = 5432
	pguser     = "postgres"
	pgpassword = "password"
	pgdbName   = "library"
)

var engine *xorm.Engine

func InitializeDB() (*Stores, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pghost, pgport, pguser, pgpassword, pgdbName)
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

func SyncStructs(db *xorm.Engine) error {
	db.Sync()
	return db.Sync2(new(User))
}
