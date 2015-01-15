package root

import (
	"app/controllers"
	"app/helper"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	//_ "github.com/go-sql-driver/mysql"
)

type UploadController struct {
	controllers.BaseController
}

func (this *UploadController) Index() {
	beego.Debug("Upload Test-------->>>>>>>>>")
	this.TplNames = "root/upload.html"
}

func (this *UploadController) Form() {
	_, rs := helper.Upload(this.Ctx.ResponseWriter, this.Ctx.Request)
	if rs == true {
		this.Redirect("/", 302)
	} else {
		fmt.Fprintf(os.Stdout, "上传失败了")
	}
	return
}

func (this *UploadController) Ajax() {
	path, rs := helper.Upload(this.Ctx.ResponseWriter, this.Ctx.Request)
	if rs == true {
		this.Ctx.WriteString(`{"code": 1, "msg": "OK", "path": "` + path + `"}`)
	} else {
		this.Ctx.WriteString(`{"code": 0, "msg": "呵呵"}`)
	}
	return
}
