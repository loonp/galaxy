package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	m "loonp/models"
	"loonp/service"
	"reflect"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	this.TplNames = "user/index.html"
}

func (this *UserController) List() {
	var (
		page     int
		pagesize int
		err      error
	)

	if page, err = strconv.Atoi(this.GetString("page")); err != nil || page < 1 {
		page = 1
	}

	pagesize = 10

	fmt.Println("start query users")
	//qs查询倒序采用字段前加 - 方式
	users, count := service.GetUserList(int64(page), int64(pagesize), "register_time")
	fmt.Println("page", page)
	this.Data["users"] = users
	this.Data["pagetotal"] = count
	this.Data["pagebar"] = m.NewPager(int64(page), int64(count), int64(pagesize), "").ToString()
	//fmt.Println("pagebar", m.NewPager(int64(page), int64(count), int64(pagesize), "").ToString())
	this.TplNames = "user/list.html"
}

func (this *UserController) Edit() {
	fmt.Println("edit user")
	userId, _ := this.GetInt64("userId")
	user := service.GetUserById(userId)
	this.Data["user"] = user
	this.TplNames = "user/edit.html"
}

func (this *UserController) Create() {
	valid := validation.Validation{}
	username := this.GetString("username")
	password := this.GetString("password")
	email := this.GetString("email")
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/manage")
	}

	valid.Email(email, "email")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
		}
	}

	fmt.Println(reflect.TypeOf(userinfo))
	//这种一种类型转换的方式
	userid := userinfo.(m.User).Id
	//var userid int64 = 1
	fmt.Println(userid)
	//userid = userinfo.Id
	//userid = strconv.ParseInt(userinfo["id"], 10, 64)
	id, _ := service.AddUser(username, password, email)
	//this.Data["id"] = id
	fmt.Println("id", id)
	//this.TplNames = "content/list.html"
	this.Ctx.Redirect(302, "/user/list")
}

func (this *UserController) Update() {
	username := this.GetString("username")
	userId, _ := this.GetInt64("userId")
	email := this.GetString("email")
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/manage")
	}

	service.UpdateUser(userId, username, email)
	this.Ctx.Redirect(302, "/user/list")
}

func (this *UserController) Del() {
	userId, _ := this.GetInt64("userId")
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/manage")
	}
	num, _ := service.DelUser(userId)
	fmt.Println(num)
	this.Ctx.Redirect(302, "/user/list")
}
