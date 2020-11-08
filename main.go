package main

import (
	"coffee-break/models"
	_ "coffee-break/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	models.Init()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		logs.Error("Error loading .env file")
		panic(err)
	}
	beego.Run(":" + os.Getenv("APP_PORT"))
}
