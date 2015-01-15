package root

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"app/controllers"
	"app/helper"
	"app/models"
)

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) Get() {
	this.Data["IsLoginP"] = true
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.DelSession("Userid")
		this.DelSession("Username")
		this.Redirect("/login", 302)
		return
	}
	this.TplNames = "root/login.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")

	password = helper.GetMD5(password)

	user := models.GetUserByName(username)
	beego.Debug("Returning ------>>>>", user)
	if username == user.Username && password == user.Password {
		this.SetSession("Userid", int64(user.Id))
		this.SetSession("Username", user.Username)
		this.Redirect("/root", 302)
	} else {
		//msg := "views/msg.html"
		this.Redirect("/login", 302)
	}
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := ck.Value
	return username == "vp" && password == "123456"
}
