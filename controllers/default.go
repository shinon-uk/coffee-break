package controllers

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"log"
)

var db *sql.DB
var connectionError error

// DB の接続情報
const (
	DriverName = "mysql"
	DataSourceName = "user:password@tcp(coffee_break_db:3306)/coffee_break"
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
	fmt.Println("Get DB Info")
	rows, err := db.Query("SELECT id, value1, value2 FROM sample")
	if err != nil {
		log.Print("error executing database query: ", err)
		panic(err)
	}
	defer rows.Close()

	var id int
	var value1 string
	var value2 string
	for rows.Next() {
		if err := rows.Scan(&id, &value1, &value2); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, value1, value2)
	}
	fmt.Println("done")

	return []string{
		value1,
		value2,
	}
}
