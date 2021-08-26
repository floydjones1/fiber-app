package data

import (
	"time"

	"xorm.io/xorm"
)

//go:generate mockery --name UserStorer --structname MockUserStore --filename user.go

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	IsDeleted bool      `json:"-"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"created"`
}

func (m *User) TableName() string {
	return "user"
}

type UserStore struct {
	db *xorm.Engine
}

var _ UserStorer = (*UserStore)(nil)

type UserStorer interface {
	GetUserByEmail(email string) (*User, bool, error)
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
