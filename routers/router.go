package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/youhao", &controllers.YouhaoController{})
}
