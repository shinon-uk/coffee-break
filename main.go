package main

import (
	"coffee-break/controllers"
	_ "coffee-break/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	result := controllers.Find()
	fmt.Println(result)
	beego.Run(":10080")
}

