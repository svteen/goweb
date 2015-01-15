package root

import (
	"strconv"

	//"github.com/astaxie/beego"

	"app/controllers"
	"app/models"
)

type PostController struct {
	controllers.BaseController
}

func (this *PostController) Get() {
	IsLogin := this.BaseController.CheckLogin()
	if IsLogin == false {
		this.Redirect("/login", 302)
		return
	}
	action := this.Input().Get("action")

	if action == "edit" {
		this.Data["IsEdit"] = true
		userid := this.GetSession("Userid").(int64)
		id := this.Input().Get("id")
		transit, _ := strconv.Atoi(id)
		post := models.GetAppointedArticle(userid, transit)
		this.Data["post"] = post
		this.TplNames = "root/post.html"
		return
	}

	if action == "del" {
		id := this.Input().Get("id")
		did, _ := strconv.Atoi(id)
		rs := models.DelArticle(did)
		if rs != 0 {
			this.Redirect("/root/post/list", 302)
		}
		return
	}

	this.Data["Username"] = this.BaseController.Data["Username"]
	this.Data["IsPost"] = true
	this.Data["IsLogined"] = this.BaseController.Data["IsLogined"]
	this.TplNames = "root/post.html"
}

func (this *PostController) Post() {
	title := this.Input().Get("title")
	excerpt := this.Input().Get("excerpt")
	origin := this.Input().Get("origin")
	content := this.Input().Get("content")

	action := this.Input().Get("action")
	var rs int64
	if action == "edit" {
		id := this.Input().Get("id")
		transit, _ := strconv.Atoi(id)
		rs = models.UpdatePost(transit, title, excerpt, origin, content)
	} else {
		rs = models.SavePost(title, excerpt, origin, content)
	}
	if rs != 0 {
		this.Redirect("/root/post/list", 302)
		return
	}

	this.Redirect("/root/post", 302)
	return
}

func (this *PostController) List() {
	this.Ctx.WriteString("List???")
}
