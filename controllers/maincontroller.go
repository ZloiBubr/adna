package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	beego.Info("MainController called!")

	c.Data["Website"] = "127.0.0.1"
	c.Data["WebsiteHTTPS"] = "127.0.0.1:8080/v1/main"
	c.Data["Email"] = "siarhei.hladkou@gmail.com"
	c.TplName = "index.tpl"

}
