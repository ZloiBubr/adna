package controllers

import (
	"github.com/astaxie/beego"
)

type SecondaryController struct {
	beego.Controller
}

func (c *SecondaryController) Get() {
	beego.Info("SecondaryController called!")

	c.Data["Website"] = "adna.com"
	c.Data["Email"] = "siarhei.hladkou@gmail.com"
	c.TplName = "secure.tpl"

}
