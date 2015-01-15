package root

import (
	"app/controllers"
	//"app/helper"
	"app/models"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type PostListController struct {
	controllers.BaseController
}

func (this *PostListController) Get() {
	IsLogin := this.BaseController.CheckLogin()
	if IsLogin == false {
		this.Redirect("/login", 302)
		return
	}
	this.Data["Username"] = this.BaseController.Data["Username"]
	this.Data["IsPostList"] = true
	this.Data["IsLogined"] = this.BaseController.Data["IsLogined"]
	//Pagination
	per := 10
	nums := models.GetPostNums()
	nums = int64(nums)
	paginator := this.SetPaginator(per, nums)
	this.Data["pages"] = paginator.Pages()
	var offset int
	if paginator.Page() == 1 {
		offset = 0
	} else {
		offset = per * (paginator.Page() - 1)
	}
	this.Data["PagingData"] = models.GetPaginationPost(per, offset)
	this.TplNames = "root/postlist.html"
}

func (this *PostListController) SetPaginator(per int, nums int64) *pagination.Paginator {
	p := pagination.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}
