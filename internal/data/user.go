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
	Email     int64
	Password  bool
	IsDeleted string
	CreatedAt time.Time
	UpdatedAt time.Time
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

func (u *UserStore) GetUserByEmail(email string) (User, error) {
	user := new(User)
	_, err := u.db.Where("email = ?", email).Get(user)
	if err != nil {
		log.Err(err).Msgf("failed to find user")
	}
	fmt.Println(err)
	return *user, err
}
