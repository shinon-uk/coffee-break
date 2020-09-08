package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetContents() {
	c.Data["json"] = privateFunction()
	c.ServeJSON()
}

// メソッド名の先頭が小文字がプライベートメソッド
func privateFunction() []string {
	return []string{
		"value1",
		"value2",
	}
}
