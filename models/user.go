package models

import "time"

type User struct {
	Id        int64
	UserName  string
	Password  string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func (m *User) TableName() string {
	return "user"
}
