package routers

import (
	"coffee-break/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.DelStaticPath("/static")
	beego.SetStaticPath("//", "dist")

	// api namespace
	apiNS := beego.NewNamespace("/api",

		// Handle user requests
		beego.NSRouter("/getContents", &controllers.MainController{}, "get:GetContents"),
	)

	beego.AddNamespace(apiNS)
}
