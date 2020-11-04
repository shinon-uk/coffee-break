package main

import (
	_ "coffee-break/routers"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

// 初期化処理
func init() {
	fmt.Println("init")
}

func main() {
	beego.Run(":10080")
}
