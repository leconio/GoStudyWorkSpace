package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello/?:id", &controllers.HelloController{},"get:SayHello")
	beego.Router("/hello/config/:env", &controllers.HelloController{},"get:SayConfig")
	beego.Router("/template", &controllers.TemplateController{})
	beego.Router("/layout", &controllers.LayoutController{})
}
