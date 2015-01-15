package controllers

import (
//"github.com/astaxie/beego"
)

type AboutController struct {
	BaseController
}

func (this *AboutController) Get() {
	this.Data["IsAbout"] = true
	this.Data["IsLogined"] = this.BaseController.Data["IsLogined"]
	this.TplNames = "about.html"
}
