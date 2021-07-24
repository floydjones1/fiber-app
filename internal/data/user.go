package data

import (
	"fmt"
	"time"

	"xorm.io/xorm"
)

type User struct {
	ID        int64
	Name      string
	DOB       time.Time
	Age       int64
	IsMarried bool
}

type UserStore struct {
	db *xorm.Engine
}
type UserStorer interface {
	GetUser(int64) error
}

func (u *UserStore) GetUser(id int64) error {
	user := new(User)
	has, err := u.db.Get(user)
	if has {
		fmt.Println("found user")
	}
	fmt.Println(err)
	return err
}
