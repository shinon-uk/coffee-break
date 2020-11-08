package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/joho/godotenv"
	"os"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		logs.Error("Error loading .env file")
		panic(err)
	}
	err = orm.RegisterDriver(
		"mysql",
		orm.DRMySQL,
	)
	if err != nil {
		logs.Error("RegisterDriver err")
		panic(err)
	}
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	port := os.Getenv("DB_PORT")
	err = orm.RegisterDataBase(
		"default",
		"mysql",
		user+":"+password+"@tcp(coffee_break_db:"+port+")/"+database,
	)
	if err != nil {
		logs.Error("RegisterDataBase err")
		panic(err)
	}
	orm.RegisterModel(new(User))
}
