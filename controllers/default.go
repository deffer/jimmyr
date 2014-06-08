package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type MainPageController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *MainPageController) Get() {
	//this.Data["Website"] = "beego.me"
	//this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "main.tpl"
}
