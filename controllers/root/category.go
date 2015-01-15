package root

import (
	"strconv"

	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"

	"app/controllers"
	"app/models"
)

type CategoryController struct {
	controllers.BaseController
}

func (this *CategoryController) Get() {
	IsLogin := this.BaseController.CheckLogin()
	if IsLogin == false {
		this.Redirect("/login", 302)
		return
	}

	this.Data["Username"] = this.BaseController.Data["Username"]
	this.Data["IsCategory"] = true
	this.Data["IsLogined"] = this.BaseController.Data["IsLogined"]

	allCate := models.GetAllCategory()
	this.Data["AllCategory"] = allCate

	action := this.Input().Get("action")
	if action == "edit" {
		this.Data["IsEdit"] = true
		id := this.Input().Get("id")
		eid, _ := strconv.Atoi(id)
		category := models.GetCategory(eid)
		this.Data["category"] = category
		this.TplNames = "root/category.html"
		return
	} else if action == "del" {
		id := this.Input().Get("id")
		did, _ := strconv.Atoi(id)
		rs := models.DelCategory(did)
		if rs != 0 {
			this.Redirect("/root/category", 302)
		}
		return
	}

	//Pagination
	per := 10
	nums := models.GetCategoryNums()
	nums = int64(nums)
	paginator := this.SetPaginator(per, nums)
	this.Data["pages"] = paginator.Pages()
	var offset int
	if paginator.Page() == 1 {
		offset = 0
	} else {
		offset = per * (paginator.Page() - 1)
	}
	this.Data["PagingData"] = models.GetPaginationCate(per, offset)

	this.TplNames = "root/category.html"
}

func (this *CategoryController) Post() {
	IsLogin := this.BaseController.CheckLogin()
	if IsLogin == false {
		this.Redirect("/login", 302)
		return
	}
	pid := this.Input().Get("pid")
	transitPid, _ := strconv.Atoi(pid)
	catename := this.Input().Get("catename")
	models.SaveCategory(transitPid, catename)
	this.Redirect("/root/category", 302)
}

func (this *CategoryController) SetPaginator(per int, nums int64) *pagination.Paginator {
	p := pagination.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}
