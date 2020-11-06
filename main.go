package main

import (
	"coffee-break/models"
	_ "coffee-break/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.Init()
}

func main() {
	beego.Run(":10080")
}
