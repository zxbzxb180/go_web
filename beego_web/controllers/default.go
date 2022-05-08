package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	website := c.GetString("website")
	c.Data["Website"] = website
	c.Data["Email"] = "616529325@qq.com"
	c.TplName = "index.tpl"
}
