package routers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"

	"adna/redirect"
	"adna/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1",
		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "127.0.0.1" {
				return true
			}
			return false
		}),
		beego.NSBefore(redirect.Https),
		beego.NSNamespace("/main",
			beego.NSRouter("/",
				&controllers.SecondaryController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

