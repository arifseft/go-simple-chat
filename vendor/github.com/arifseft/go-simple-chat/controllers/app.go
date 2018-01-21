package controllers

import (
	"github.com/astaxie/beego"
)


type baseController struct {
	beego.Controller
}

type AppController struct {
	baseController
}

func (this *AppController) Get() {
	this.TplName = "welcome.html"
}

func (this *AppController) Join() {
	uname := this.GetString("uname")
	group := this.GetString("group")

	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	this.Redirect("/chat?uname="+uname+"&group="+group, 302)
	
	return
}
