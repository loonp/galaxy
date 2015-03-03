package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/loonp/galaxy/lib"
	m "github.com/loonp/galaxy/models"
	"github.com/loonp/galaxy/service"
	"strconv"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index() {
	var (
		page     int
		pagesize int
		err      error
	)

	if page, err = strconv.Atoi(this.GetString("page")); err != nil || page < 1 {
		page = 1
	}

	pagesize = 10

	fmt.Println("start index contents")
	contents, count := service.GetContentList(int64(page), int64(pagesize), "-update_time", "-create_time")

	channelId, _ := this.GetInt64("channelId")
	if channelId > 0 {
		contents, count = service.GetContentListByChannelId(int64(page), int64(pagesize), int64(channelId), "-update_time", "-create_time")
	}

	fmt.Println("tt", contents, count)
	// fmt.Println("page", page)

	this.Data["contents"] = contents
	this.Data["pagetotal"] = count

	this.Data["pagebar"] = m.NewPager(int64(page), int64(count), int64(pagesize), "").ToStringIndex(channelId)

	//fmt.Println("pagebar", m.NewPager(int64(page), int64(count), int64(pagesize), "").ToStringIndex())

	channels, _ := service.GetAllChannel()
	this.Data["channels"] = channels

	fmt.Println(channels)

	this.TplNames = "index.tpl"
}

func (this *IndexController) Detail() {
	contentId, _ := this.GetInt64("contentId")
	content := service.GetContentById(contentId)
	this.Data["content"] = content

	channels, _ := service.GetAllChannel()
	this.Data["channels"] = channels

	this.TplNames = "detail.tpl"

}

func (this *IndexController) Test() {
	fmt.Println(lib.HashDatePath("123"))
	fmt.Println(lib.HashDateName())

	/***
	测试json格式
	*/
	type Ret struct {
		Error int64  `json:"error"`
		Url   string `json:"url"`
	}
	this.Data["json"] = Ret{Error: 1, Url: "baidu"}
	this.ServeJson()

}

func (this *IndexController) Add() {
	fmt.Println("test")

}
