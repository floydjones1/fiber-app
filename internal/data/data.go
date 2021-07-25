package data

import (
	"fmt"

	"github.com/floydjones1/fiber-app/config"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type Stores struct {
	DB *xorm.Engine
	UserStore
	BookStore
}

var engine *xorm.Engine

func InitializeDB(dbConf config.Database) (*Stores, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.DatabaseName)
	engine, err := xorm.NewEngine("postgres", connectionInfo)
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
