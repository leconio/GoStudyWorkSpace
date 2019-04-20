package controllers

import "github.com/astaxie/beego"

type LayoutController struct {
	beego.Controller
}

func (this *LayoutController) Get()  {
	this.Layout = "base/layout.html"
	this.TplName = "layout_test.html"

}
