package controllers

import (
	"github.com/astaxie/beego"
)

type YouhaoController struct {
	beego.Controller
}

func (c *YouhaoController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
