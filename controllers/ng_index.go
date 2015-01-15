package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"

	/*"app/helper"*/
	"app/models"
)

const (
	Per = 10
)

type NgController struct {
	BaseController
}

func (this *NgController) Get() {
	beego.Debug("Angularjs page")
	id := this.Input().Get("id")
	var pid int
	if id == "" {
		pid = 0
	} else {
		pid, _ = strconv.Atoi(id)
	}

	nums := models.GetPostNums()
	nums = int64(nums)

	this.SetPaginator(Per, nums)
	beego.Debug("pid---------<", pid)
	beego.Debug("nums--------->", nums)
	this.TplNames = "ng-index.html"
}

func (this *NgController) SetPaginator(per int, nums int64) *pagination.Paginator {
	p := pagination.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}
