package controllers

import (
	"coffee-break/models"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	baseController
	repository models.UserRepository
}

type User struct {
	Name     string `form:"username"`
	Password string `form:"password"`
}

func (c *UserController) FetchAllUser() {
	users, err := c.repository.FindAll()
	if err != nil {
		logs.Error("FindAll err")
		panic(err)
	}
	c.Data["json"] = users
	c.ServeJSON()
}

func (c *UserController) AddUser() {
	// バリデーションチェック
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		logs.Error("ParseForm err")
		panic(err)
	}
	if u.Name == "" {
		panic("ユーザー名が入力されてません")
	}
	if u.Password == "" {
		panic("パスワードが入力されてません")
	}

	// 保存処理
	id, err := c.repository.Save(
		&models.User{
			UserName: c.GetString("username"),
			Password: c.GetString("password"),
		},
	)
	if err != nil {
		logs.Error("Insert err")
		panic(err)
	}

	// 確認
	users, err := c.repository.FindById(id)
	if err != nil {
		logs.Error("FindAll err")
		panic(err)
	}
	logs.Info(users)

	c.ServeJSON()
}
