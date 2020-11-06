package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

const (
	DriverName     = "mysql"
	DataSourceName = "user:password@tcp(coffee_break_db:3306)/coffee_break"
)

func Init() {
	err := orm.RegisterDriver(DriverName, orm.DRMySQL)
	if err != nil {
		logs.Error(err)
	}
	err = orm.RegisterDataBase("default", DriverName, DataSourceName)
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModel(new(Sample))
}
