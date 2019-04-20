package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
)

type TemplateController struct {
	beego.Controller
}

type User struct {
	Id    int    `form:"-"`
	Name  string `form:"username"`
	Age   string `form:"age"`
	Email string
}

func (this *TemplateController) Get() {
	this.Data["name"] = "lecon"
	this.Data["email"] = "leconio@outlook.com"
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	u := this.GetSession("userForm")
	if u != nil {
		this.Data["userForm"] = &u
	}
	this.TplName = "template.html"
}

func (this *TemplateController) Post() {
	u := User{}
	if err := this.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	this.SetSession("userForm",u)
	this.Data["json"] = u
	this.ServeJSON()
}
