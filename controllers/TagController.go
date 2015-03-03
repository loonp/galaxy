package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	m "github.com/loonp/galaxy/models"
	"github.com/loonp/galaxy/service"
	// "reflect"
	// "strconv"
)

type TagController struct {
	beego.Controller
}

func (this *TagController) Create() {
	name := this.GetString("name")
	tag, err := service.GetTagByName(name)

	if err != orm.ErrNoRows {
		this.Data["json"] = tag
	} else {
		var id int64
		id, err = service.AddTag(name)
		this.Data["json"] = m.Tag{Id: id, Name: name}
	}
	this.ServeJson()
}
