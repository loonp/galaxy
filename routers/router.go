package routers

import (
	"github.com/astaxie/beego"
	"github.com/loonp/galaxy/controllers"
	m "github.com/loonp/galaxy/models"
	"reflect"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	beego.Router("/", &controllers.IndexController{}, "get,post:Index")
	beego.Router("/detail", &controllers.IndexController{}, "get,post:Detail")
	beego.Router("/test", &controllers.IndexController{}, "get:Test")

	beego.Router("/manage", &controllers.ManageController{})
	//beego.Router("/manage/login", &controllers.ManageController{}, "get:Get")
	beego.Router("/manage/login", &controllers.ManageController{}, "get,post:Login")
	beego.Router("/manage/logout", &controllers.ManageController{}, "get,post:Logout")

	beego.Router("/content/del", &controllers.ContentController{}, "get,post:Del")
	beego.Router("/content/edit", &controllers.ContentController{}, "get,post:Edit")
	beego.Router("/content/create", &controllers.ContentController{}, "get,post:Create")
	beego.Router("/content/addindex", &controllers.ContentController{}, "get:AddIndex")
	beego.Router("/content/list", &controllers.ContentController{}, "get,post:List")
	beego.Router("/content/update", &controllers.ContentController{}, "get,post:Update")
	beego.Router("/content/upload", &controllers.ContentController{}, "post:Upload")

	//channelrouter
	beego.Router("/channel/list", &controllers.ChannelController{}, "get,post:List")
	beego.Router("/channel/addindex", &controllers.ChannelController{}, "get,post:AddIndex")
	beego.Router("/channel/delete", &controllers.ChannelController{}, "get,post:DeleteChannel")
	beego.Router("/channel/create", &controllers.ChannelController{}, "get,post:AddChannel")
	beego.Router("/channel/toupdatechannel", &controllers.ChannelController{}, "get,post:ToUpdateChannel")
	beego.Router("/channel/update", &controllers.ChannelController{}, "get,post:UpdateChannel")

	//media router
	beego.Router("/media/list", &controllers.MediaController{}, "get,post:List")
	beego.Router("/media/addindex", &controllers.MediaController{}, "get,post:AddIndex")
	beego.Router("/media/upload", &controllers.MediaController{}, "post:Upload")
	beego.Router("/media/uploadin", &controllers.MediaController{}, "post:UploadIn")

	//user router
	beego.Router("/user/list", &controllers.UserController{}, "get,post:List")
	beego.Router("/user/index", &controllers.UserController{}, "get,post:Index")
	beego.Router("/user/create", &controllers.UserController{}, "get,post:Create")
	beego.Router("/user/edit", &controllers.UserController{}, "get,post:Edit")
	beego.Router("/user/update", &controllers.UserController{}, "get,post:Update")
	beego.Router("/user/del", &controllers.UserController{}, "get,post:Del")

	//tag router
	beego.Router("/tag/create", &controllers.TagController{}, "get,post:Create")

	beego.AddFuncMap("getUsernameById", GetUsernameById)
	beego.AddFuncMap("getChannelTypeById", GetChannelTypeById)
	beego.AddFuncMap("getChannelNameById", GetChannelNameById)
	beego.AddFuncMap("getTagNameById", GetTagNameById)
	beego.AddFuncMap("getTagNamesByCid", GetTagNamesByCid)
	beego.AddFuncMap("typeOf", TypeOf)

}

/**
模板函数，这个暂时的解决联表查询多字段的问题，会影响数据库的压力,后期一定要修改
*/
func GetUsernameById(id int64) string {
	user := m.GetUserById(id)
	return user.Username
}

func GetChannelTypeById(id int64) string {
	switch id {
	case 0:
		return "图文"
	case 1:
		return "媒体"
	case 2:
		return "内容型分类"
	case 3:
		return "图片加描述"
	default:
		return "未知"
	}
}

func GetChannelNameById(id int64) string {
	channel, _ := m.GetChannelById(id)
	return channel.ChannelName
}

func TypeOf(object interface{}) reflect.Type {
	return reflect.TypeOf(object)
}

func GetTagNameById(id int64) string {
	tag, _ := m.GetTagById(id)
	return tag.Name
}

func GetTagNamesByCid(cid int64) string {
	contentTags, _ := m.GetContentTagsByCid(cid)
	var ret string
	for i, contentTag := range contentTags {
		if i == 0 {
			tag, _ := m.GetTagById(contentTag["Tid"].(int64))
			ret = tag.Name
		} else {
			tag, _ := m.GetTagById(contentTag["Tid"].(int64))
			ret = ret + " " + tag.Name
		}
	}
	return ret
}
