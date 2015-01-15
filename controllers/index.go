package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"

	"app/helper"
	"app/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["password"] = helper.GetMD5("123456")
	this.Data["IsLogined"] = this.BaseController.Data["IsLogined"]
	this.Data["Userid"] = this.GetSession("Userid")
	this.Data["Username"] = this.GetSession("Username")
	this.Data["IsHome"] = true
	//Pagination
	per := 10
	nums := models.GetPostNums()
	nums = int64(nums)
	paginator := this.SetPaginator(per, nums)
	this.Data["pages"] = paginator.Pages()
	beego.Debug("SQL------------>", 1)
	var offset int
	if paginator.Page() == 1 {
		offset = 0
	} else {
		offset = per * (paginator.Page() - 1)
	}
	this.Data["PagingData"] = models.GetPaginationPost(per, offset)

	/*s := [][]string{{"a", "b", "sssss"}, {"c", "d", "ssss"}, {"c", "d", "sss"}}
	helper.Write2Excel("./static/files/abc.xls", s)*/
	/*msg := "views/msg.html"
	this.Data["Icontent"], _ = helper.RenderPartial(msg, this.Data)*/
	this.TplNames = "index.html"
}

func (this *MainController) SetPaginator(per int, nums int64) *pagination.Paginator {
	p := pagination.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}
