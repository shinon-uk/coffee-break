package main

import (
	_ "coffee-break/routers"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var connectionError error

// DB の接続情報
const (
	DriverName = "mysql"
	DataSourceName = "user:password@tcp(coffee_break_db:3306)/coffee_break"
)

// 初期化処理
func init() {
	db, err := sql.Open(DriverName, DataSourceName)
	if err != nil {
		fmt.Println("open err")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ping err")
		panic(err)
	}

	fmt.Println("ok")
	defer db.Close()
}

func main() {
	beego.Run(":10080")
}
