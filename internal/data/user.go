package data

import (
	"time"

	"xorm.io/xorm"
)

type User struct {
	Id        int64
	Name      string
	Email     string
	Password  string
	IsDeleted bool
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"created"`
}

func (m *User) TableName() string {
	return "user"
}

type UserStore struct {
	db *xorm.Engine
}
type UserStorer interface {
	GetUser(int64) (User, error)
	InsertUser(user *User) error
}

func (u *UserStore) GetUserByEmail(email string) (*User, bool, error) {
	user := new(User)
	found, err := u.db.Where("email = ?", email).Get(user)
	if err != nil {
		return nil, found, err
	}
	return user, found, err
}

func (u *UserStore) InsertUser(user *User) error {
	_, err := u.db.Insert(user)
	return err
}
