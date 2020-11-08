package routers

import (
	"coffee-break/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.DelStaticPath("/static")
	beego.SetStaticPath("//", "static/dist")
	beego.Router("/api/fetchAllUser", &controllers.UserController{}, "get:FetchAllUser")
	beego.Router("/api/addUser", &controllers.UserController{}, "post:AddUser")
}
