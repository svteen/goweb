package root

import (
	"app/controllers"
	// "app/models"
	//"github.com/astaxie/beego"
)

type HomeController struct {
	controllers.BaseController
}

func (this *HomeController) Get() {
	IsLogin := this.BaseController.CheckLogin()

	if IsLogin == false {
		this.Redirect("/login", 302)
		return
	}
	this.Data["Username"] = this.BaseController.Data["Username"]
	this.Data["IsHome"] = true
	this.Data["IsLogined"] = this.BaseController.Data["IsLogined"]
	this.TplNames = "root/home.html"
}
