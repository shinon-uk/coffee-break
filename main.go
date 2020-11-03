package main

import (
	_ "coffee-break/routers"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

	fmt.Println("Get DB Info")
	rows, err := db.Query("SELECT id, value1, value2 FROM sample")
	if err != nil {
		log.Print("error executing database query: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var value1 string
		var value2 string
		if err := rows.Scan(&id, &value1, &value2); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, value1, value2)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("select ok")

	result, err := db.Exec("UPDATE sample SET value1 = ?, value2 = ? WHERE id = ?", "aaa", "aaa", 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	fmt.Println("update ok")

	defer db.Close()
}

func main() {
	beego.Run(":10080")
}
