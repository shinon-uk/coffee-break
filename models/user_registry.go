package models

import (
	"github.com/astaxie/beego/orm"
)

type UserRepository struct {
}

func (r *UserRepository) FindAll() ([]User, error) {
	o := orm.NewOrm()
	var users []User
	_, err := o.QueryTable("user").All(&users, "UserName")
	return users, err
}

func (r *UserRepository) FindById(id int64) (*User, error) {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	return &user, err
}

func (r *UserRepository) Save(u *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(u)
	return id, err
}

func (r *UserRepository) Delete(u *User) error {
	o := orm.NewOrm()
	_, err := o.Delete(u)
	return err
}
