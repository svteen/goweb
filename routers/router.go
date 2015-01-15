package routers

import (
	"app/controllers"
	"app/controllers/root"
	"github.com/astaxie/beego"
)

func init() {
	//frontend
	beego.Router("/", &controllers.MainController{})
	beego.Router("/lottery", &controllers.LotteryController{})
	beego.Router("/about", &controllers.AboutController{})
	//frontend angular
	beego.Router("/index", &controllers.NgController{})

	//backend
	beego.Router("/login", &root.LoginController{})
	beego.AutoRouter(&root.UploadController{})
	beego.Router("/root", &root.HomeController{})
	beego.Router("/root/category", &root.CategoryController{})
	beego.Router("/root/post", &root.PostController{})
	beego.Router("/root/post/list", &root.PostListController{})
	beego.Router("/root/account", &root.AccountController{})
	beego.Router("/article/:id:int", &controllers.ArticleController{})
}
