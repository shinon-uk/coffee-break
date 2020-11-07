package controllers

import (
	"coffee-break/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	baseController
	repository models.UserRepository
}

type User struct {
	Name     string `validate:"required" form:"username"`
	Password string `validate:"required" form:"password"`
}

func (c *UserController) FetchAllUser() {
	//users, err := c.repository.FindAll()
	o := orm.NewOrm()
	var users []models.User
	o.QueryTable("user").All(&users, "UserName")

	//if err != nil {
	//	logs.Error("FindAll err")
	//	panic(err)
	//}
	logs.Info(users)
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
