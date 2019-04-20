package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (this *HelloController) SayHello() {
	fmt.Println(this.Ctx.Input.Param(":id"))
	this.Ctx.WriteString("HelloWorld")
}

func (this *HelloController) SayConfig() {
	env := this.Ctx.Input.Param(":env")
	resp := beego.AppConfig.String(env+"::mysqluser") + "\n" +
		beego.AppConfig.String(env+"::mysqlpass") + "\n" +
		beego.AppConfig.String("mysqlurls") + "\n" +
		beego.AppConfig.String("mysqldb")
	this.Ctx.WriteString(resp)
}
