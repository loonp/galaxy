package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"galaxy/models"
	_ "github.com/loonp/galaxy/routers"
)

func init() {
	//models.RegisterDB()
}

func main() {
	orm.RunSyncdb("default", false, true)
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/fonts", "static/fonts")

	beego.SetStaticPath("/mm", "mm")

	beego.EnableAdmin = true
	beego.Run()
}
