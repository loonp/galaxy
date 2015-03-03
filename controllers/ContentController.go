package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/loonp/galaxy/lib"
	m "github.com/loonp/galaxy/models"
	"github.com/loonp/galaxy/service"
	"reflect"
	"strconv"
	"strings"
)

type ContentController struct {
	beego.Controller
}

func (this *ContentController) Get() {

}

func (this *ContentController) Create() {
	contentTitle := this.GetString("contentTitle")
	contentMain := this.GetString("contentMain")
	contentSummary := this.GetString("contentSummary")
	fmt.Println(contentMain)
	channelId, _ := this.GetInt64("channelId")
	fmt.Println(contentTitle)
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/manage")
	}
	fmt.Println(reflect.TypeOf(userinfo))
	//这种一种类型转换的方式
	userid := userinfo.(m.User).Id

	//var userid int64 = 1
	fmt.Println(userid)
	//userid = userinfo.Id
	//userid = strconv.ParseInt(userinfo["id"], 10, 64)
	picurl := this.GetString("picurl")
	id, err := service.AddContent(channelId, contentTitle, contentSummary, contentMain, userid, picurl)
	//this.Data["id"] = id
	fmt.Println("id", id)
	fmt.Println(err)
	tags := this.GetString("tags")
	tagstr := strings.Split(tags, ",")
	for _, s := range tagstr {
		tid, _ := service.AddTagWithRead(s)
		service.AddContentTag(tid, id)
	}

	//this.TplNames = "content/list.html"
	this.Ctx.Redirect(302, "/content/list")
}

func (this *ContentController) Update() {
	contentTitle := this.GetString("contentTitle")
	contentMain := this.GetString("contentMain")
	contentSummary := this.GetString("contentSummary")
	fmt.Println(contentMain)
	channelId, _ := this.GetInt64("channelId")
	createId, _ := this.GetInt64("createId")
	id, _ := this.GetInt64("contentId")
	fmt.Println(contentTitle)
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/manage")
	}
	fmt.Println(reflect.TypeOf(userinfo))
	//这种一种类型转换的方式
	userid := userinfo.(m.User).Id
	//var userid int64 = 1
	fmt.Println(userid)
	//userid = userinfo.Id
	//userid = strconv.ParseInt(userinfo["id"], 10, 64)
	picurl := this.GetString("picurl")
	service.UpdateContent(id, createId, userid, channelId, contentTitle, contentSummary, contentMain, picurl)

	tags := this.GetString("tags")
	tagstr := strings.Split(tags, ",")
	for _, s := range tagstr {
		tid, _ := service.AddTagWithRead(s)
		service.AddContentTag(tid, id)
	}

	this.Ctx.Redirect(302, "/content/list")
}

func (this *ContentController) AddIndex() {
	fmt.Printf("addIndex.....")
	channels, _ := service.GetAllChannel()
	this.Data["channels"] = channels
	fmt.Println(channels)
	this.TplNames = "content/index.html"
}

func (this *ContentController) Edit() {
	fmt.Println("edit content")
	contentId, _ := this.GetInt64("contentId")
	content := service.GetContentById(contentId)
	this.Data["content"] = content
	//fmt.Println(reflect.TypeOf(content.ContentChannelId))
	channels, _ := service.GetAllChannel()
	this.Data["channels"] = channels

	contentTags, _ := service.GetContentTagsByCid(contentId)
	this.Data["contentTags"] = contentTags
	fmt.Println(channels)
	this.TplNames = "content/edit.html"
}

func (this *ContentController) List() {
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
	channelId, _ := this.GetInt("channelId")

	var contents []orm.Params
	var count int64
	fmt.Println("channelId ----", channelId)
	if channelId > 0 {
		contents, count = service.GetContentListByChannelId(int64(page), int64(pagesize), int64(channelId), "-update_time", "-create_time")
	} else {
		contents, count = service.GetContentList(int64(page), int64(pagesize), "-update_time", "-create_time")

	}
	//qs查询倒序采用字段前加 - 方式

	fmt.Println("page", page)
	this.Data["contents"] = contents
	this.Data["pagetotal"] = count
	this.Data["pagebar"] = m.NewPager(int64(page), int64(count), int64(pagesize), "").ToStringSearch(int64(channelId))
	this.Data["schannelId"] = channelId
	//fmt.Println("pagebar", m.NewPager(int64(page), int64(count), int64(pagesize), "").ToString())
	this.TplNames = "content/list.html"
}

func (this *ContentController) Del() {
	contentId, _ := this.GetInt64("contentId")
	fmt.Println("contentId ===", contentId)
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/manage")
	}
	num, _ := service.DelContent(contentId)
	fmt.Println(num)
	this.Ctx.Redirect(302, "/content/list")
}

func (this *ContentController) Upload() {
	//_, file, _ := this.GetFile("imgFile")
	_, header, _ := this.GetFile("imgFile")
	ext := strings.ToLower(header.Filename[strings.LastIndex(header.Filename, "."):])
	fmt.Println(ext)
	path := lib.HashDatePath(lib.HashDateName()) + ext
	this.SaveToFile("imgFile", path)
	type Ret struct {
		Error int    `json:"error"`
		Url   string `json:"url"`
	}
	this.Data["json"] = Ret{Error: 0, Url: beego.AppConfig.String("img_server_url") + "/" + path}
	this.ServeJson()
	//this.Ctx.WriteString("{\"error\": 0, \"url\": \"" + beego.AppConfig.String("img_server_url") + "/" + path + "\"}")
}
