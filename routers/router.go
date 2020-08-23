package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.DelStaticPath("/static")
	beego.SetStaticPath("//", "dist")
}
