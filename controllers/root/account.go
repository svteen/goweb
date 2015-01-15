package root

import (
	"app/controllers"
	"app/models"
	"github.com/astaxie/beego"
)

type AccountController struct {
	controllers.BaseController
}

func (this *AccountController) Get() {
	IsLogin := this.BaseController.CheckLogin()

	if IsLogin == false {
		this.Redirect("/login", 302)
		return
	}
	id := this.GetSession("Userid").(int64)
	userInfo := models.GetUserInfo(id)
	beego.Debug("Debug..Userinfo...........................", userInfo)
	this.Data["Nickname"] = userInfo.Nicknanme
	this.Data["Email"] = userInfo.Email
	this.Data["Avatar"] = userInfo.Avatar
	this.Data["Username"] = this.BaseController.Data["Username"]
	this.Data["IsHome"] = true
	this.Data["IsLogined"] = this.BaseController.Data["IsLogined"]
	this.TplNames = "root/account.html"
}

func (this *AccountController) Post() {
	IsLogin := this.BaseController.CheckLogin()

	if IsLogin == false {
		this.Redirect("/login", 302)
		return
	}
	id := this.GetSession("Userid").(int64)
	nickname := this.Input().Get("nickname")
	email := this.Input().Get("email")
	avatar := this.Input().Get("avatar")
	rs := models.UpdateUserInfo(id, nickname, email, avatar)
	if rs == true {
		this.Redirect("/root/account", 302)
	}
	this.Redirect("/root", 302)
}
