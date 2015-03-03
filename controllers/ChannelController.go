package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	m "github.com/loonp/galaxy/models"
	"github.com/loonp/galaxy/service"
	"strconv"
)

type ChannelController struct {
	beego.Controller
}

func (this *ChannelController) Get() {

}

func (this *ChannelController) List() {
	var (
		page     int
		pagesize int
		err      error
	)

	if page, err = strconv.Atoi(this.GetString("page")); err != nil || page < 1 {
		page = 1
	}

	pagesize = 10

	fmt.Println("start query contents")
	//qs查询倒序采用字段前加 - 方式
	channels, count := service.GetChannelList(int64(page), int64(pagesize), "-create_time")
	fmt.Println("page", page)
	this.Data["channels"] = channels
	this.Data["pagetotal"] = count
	this.Data["pagebar"] = m.NewPager(int64(page), int64(count), int64(pagesize), "").ToString()
	this.TplNames = "channel/list.html"

}

func (this *ChannelController) GetChannelById() {

}

func (this *ChannelController) AddIndex() {
	fmt.Printf("addIndex.....")
	this.TplNames = "channel/createindex.html"
}

func (this *ChannelController) AddChannel() {
	channelName := this.GetString("channelName")
	tmp, _ := this.GetInt("channelType")
	channelType := int8(tmp)
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/manage")
	}
	//这种一种类型转换的方式
	userid := userinfo.(m.User).Id
	id, err := service.AddChannel(channelName, channelType, userid)
	fmt.Println(id, err)
	this.Ctx.Redirect(302, "/channel/list")
}

func (this *ChannelController) UpdateChannel() {
	channelName := this.GetString("channelName")
	channelId, _ := this.GetInt64("channelId")
	createId, _ := this.GetInt64("createId")
	channelType, _ := this.GetInt64("channelType")
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/manage")
	}
	//这种一种类型转换的方式
	userid := userinfo.(m.User).Id
	//var userid int64 = 1
	fmt.Println(userid)
	//userid = userinfo.Id
	//userid = strconv.ParseInt(userinfo["id"], 10, 64)
	//id := service.AddContent(channelId, contentTitle, contentMain, userid)
	service.UpdateChannel(channelId, channelName, createId, int8(channelType), 0, 0)
	//this.Data["id"] = id
	//this.TplNames = "content/list.html"
	this.Ctx.Redirect(302, "/channel/list")
}

func (this *ChannelController) DeleteChannel() {
	channelId, _ := this.GetInt64("channelId")
	fmt.Println("channelId ===", channelId)
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/manage")
	}
	num, _ := service.DelChannel(channelId)
	fmt.Println(num)
	this.Ctx.Redirect(302, "/channel/list")
}

func (this *ChannelController) ToUpdateChannel() {
	fmt.Println("123")
	var channel m.Channel
	var err error
	tmp, _ := this.GetInt("channelId")
	channelId := int64(tmp)
	channel, err = service.GetChannelById(channelId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(channel)
	this.Data["editchannel"] = channel
	this.TplNames = "channel/edit.html"
}
