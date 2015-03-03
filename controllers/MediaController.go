package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/loonp/galaxy/lib"
	"strings"
)

type MediaController struct {
	beego.Controller
}

func (this *MediaController) List() {
	this.TplNames = "media/list.html"
}

func (this *MediaController) AddIndex() {
	fmt.Printf("addIndex.....")
	this.TplNames = "media/createindex.html"
}

func (this *MediaController) Upload() {
	_, file, _ := this.GetFile("imgFile")
	fmt.Println(file.Filename)
	this.SaveToFile("imgFile", lib.HashDatePath(lib.HashDateName())+".jpg")
	this.Ctx.Redirect(302, "/media/list")
}

func (this *MediaController) UploadIn() {
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
