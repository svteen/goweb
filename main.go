package main

import (
	_ "app/routers"
	//"fmt"
	_ "app/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:1111@/sf2?charset=utf8")
}

func main() {
	orm.Debug = true // 开启ORM调试
	//orm.RunSyncdb("default", false, true) //自动建表功能
	/*o := orm.NewOrm()
	o.Using("default")*/
	beego.SessionOn = true
	beego.Run()
}
