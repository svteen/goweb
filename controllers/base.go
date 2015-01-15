package controllers

import (
	// "app/models"
	"github.com/astaxie/beego"
)

var (
	sess_username string
	sess_uid      int64
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	sess_uid, _ = this.GetSession("Userid").(int64)
	sess_username, _ = this.GetSession("Username").(string)
	if sess_uid != 0 {
		beego.Debug("Debug")
		this.Data["Userid"] = sess_uid
		this.Data["Username"] = sess_username
		this.Data["IsLogined"] = true
	} else {
		this.Data["Userid"] = 0
		this.Data["Username"] = ""
	}
}

func (this *BaseController) CheckLogin() bool {
	if sess_uid == 0 {
		return false
	}
	return true
}
