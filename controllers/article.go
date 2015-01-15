package controllers

import (
	"app/models"
	//"github.com/astaxie/beego"
	"strconv"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Get() {
	id := this.Ctx.Input.Param(":id")
	transit, _ := strconv.Atoi(id)
	this.Data["Article"] = models.GetArticle(transit)
	this.TplNames = "article.html"
}
