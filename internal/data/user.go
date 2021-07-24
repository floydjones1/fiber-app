package data

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
)

type User struct {
	Id        int64
	Name      string
	Dob       time.Time
	Age       int64
	IsMarried bool
	Sexuality string
}

func (m *User) TableName() string {
	return "user"
}

type UserStore struct {
	db *xorm.Engine
}
type UserStorer interface {
	GetUser(int64) (User, error)
}

func (u *UserStore) GetUser(id int64) (User, error) {
	user := new(User)
	_, err := u.db.Get(user)
	if err != nil {
		log.Err(err).Msgf("failed to find user")
	}
	fmt.Println(err)
	return *user, err
}
