package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/loonp/galaxy/lib"
	"github.com/loonp/galaxy/service"
	// m "github.com/loonp/galaxy/models"
)

type ManageController struct {
	beego.Controller
}

func (this *ManageController) Get() {

	this.Data["Website"] = "contentcontroller"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "manage/index.html"

}

func (this *ManageController) Login() {
	fmt.Println("login ")

	userinfo := this.GetSession("userinfo")
	fmt.Println("session ==", userinfo)
	if userinfo != nil {
		this.Data["currentuser"] = userinfo
		this.TplNames = "manage/manage.html"

	} else {

		username := this.GetString("username")
		password := this.GetString("password")

		fmt.Println(username, password)
		fmt.Printf(lib.Pwdhash(password))

		user, err := service.CheckLogin(username, password)
		//fmt.Println(err)
		if err == nil {
			fmt.Println("user   === ", user)
			this.SetSession("userinfo", user)
			userinfo := this.GetSession("userinfo")
			fmt.Println(userinfo)
			this.Data["currentuser"] = userinfo
			this.TplNames = "manage/manage.html"
		} else {
			fmt.Println(err)
			this.TplNames = "manage/index.html"
		}

	}

}
func (this *ManageController) Logout() {
	this.DelSession("userinfo")
	this.Ctx.Redirect(302, "/manage/login")
}
