package controllers

import (
	"coffee-break/models"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	baseController
}

func (c *MainController) GetContents() {
	sample := models.Sample{Id: 1}
	err := c.o.Read(&sample)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	logs.Info(sample)
	c.Data["json"] = []string{
		sample.Value1,
		sample.Value2,
	}
	c.ServeJSON()
}
