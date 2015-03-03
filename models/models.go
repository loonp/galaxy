package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Content struct {
	Id               int64
	ContentChannelId int64
	ContentSummary   string `orm:"null;size(500)"`
	ContentTitle     string `form:"ContentTitle"`
	ContentMain      string `orm:"size(10000)" form:"ContentMain"`
	CreateId         int64
	CreateTime       time.Time `orm:"null"`
	IsDelete         bool      `orm:"null"`
	ContentStatus    bool      `orm:"null"`
	DeleteTime       time.Time `orm:"null"`
	UpdateId         int64     `orm:"null"`
	UpdateTime       time.Time `orm:"null"`
	MediaPath        string
}

type Channel struct {
	Id              int64
	ChannelName     string
	ChannelType     int8
	ChannelLevel    int32     `orm:"null"`
	ChannelParentId int64     `orm:"null"`
	CreateTime      time.Time `orm:"null"`
	CreateId        int64
}

type User struct {
	Id           int64
	Username     string `orm:"unique"`
	Password     string
	Email        string
	RegisterTime time.Time
}

type Media struct {
	Id          int64
	MediaName   string
	MediaPath   string
	MediaDesc   string
	ChannelId   int64
	ChannelType int8
	CreateId    int64
	CreateTime  time.Time
}

type Tag struct {
	Id   int64  `json:"id"`
	Name string `orm:"unique" json:"name"`
}

type ContentTag struct {
	Id  int64
	Cid int64
	Tid int64
}

//多字段唯一键
func (contentTag *ContentTag) TableUnique() [][]string {
	return [][]string{
		[]string{"Cid", "Tid"},
	}

}

func init() {
	// orm.RegisterModel(new(Content))
	// orm.RegisterModel(new(User))
	// orm.RegisterModel(new(Channel))
	// orm.RegisterModel(new(Media))
	// orm.RegisterModel(new(Tag))
	// orm.RegisterModel(new(ContentTag))
	RegisterDB()

}

func RegisterDB() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dsn)

}

/***
 user
***/
func GetUserByUsername(username string) (user User) {
	user = User{Username: username}
	o := orm.NewOrm()
	o.Read(&user, "Username")
	return user
}

func GetUserById(id int64) (user User) {
	user = User{Id: id}
	o := orm.NewOrm()
	o.Read(&user, "Id")
	return user
}

func AddUser(user User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&user)
	return id, err
}

func UpdateUser(user User) (id int64, err error) {
	o := orm.NewOrm()
	fmt.Println(user)
	id, err = o.Update(&user)
	return id, err
}

func DelUser(userId int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&User{Id: userId})
	return num, err
}

/**
	content
***/
func AddContent(content Content) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&content)
	return id, err
}

func GetContentById(id int64) (content Content) {
	content = Content{Id: id}
	o := orm.NewOrm()
	o.Read(&content, "Id")
	return content
}

func DelContent(contentId int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Content{Id: contentId})
	return num, err
}

func UpdateContent(content Content) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Update(&content)
	return id, err
}

/**
media
*/

func AddMedia(mediaName string, mediaPath string, mediaDesc string, channelId int64, channelType int8, createId int64) (id int64, err error) {
	o := orm.NewOrm()
	media := Media{MediaName: mediaName, MediaPath: mediaPath, MediaDesc: mediaDesc, ChannelId: channelId, ChannelType: channelType, CreateId: createId}
	media.CreateTime = time.Now()
	id, err = o.Insert(&media)
	return id, err
}

/***
channel

*/
func GetChannelById(channelId int64) (channel Channel, err error) {
	o := orm.NewOrm()
	err = o.Raw("select * from channel where id = ?", channelId).QueryRow(&channel)
	if err != nil {
		fmt.Println(err)
	}
	return channel, err
}

func UpdateChannel(channel Channel) {
	o := orm.NewOrm()
	o.Update(&channel)
}

func DelChannel(channelId int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Channel{Id: channelId})
	return num, err
}

func AddChannel(channelName string, channelType int8, createId int64) (id int64, err error) {
	o := orm.NewOrm()
	channel := Channel{ChannelName: channelName, ChannelType: channelType, CreateId: createId}
	channel.CreateTime = time.Now()
	id, err = o.Insert(&channel)
	return id, err
}

/**
 *
 */
func AddTag(name string) (id int64, err error) {
	o := orm.NewOrm()
	tag := Tag{Name: name}
	id, err = o.Insert(&tag)
	return id, err
}

func GetTagByName(name string) (tag Tag, err error) {
	tag = Tag{Name: name}
	o := orm.NewOrm()
	err = o.Read(&tag, "name")
	return tag, err
}

func GetTagById(id int64) (tag Tag, err error) {
	tag = Tag{Id: id}
	o := orm.NewOrm()
	err = o.Read(&tag, "id")
	return tag, err
}

/**
**/
func AddContentTag(tid int64, cid int64) (id int64, err error) {
	o := orm.NewOrm()
	contentTag := ContentTag{Tid: tid, Cid: cid}
	id, err = o.Insert(&contentTag)
	return id, err
}

func GetContentTagsByCid(cid int64) (contentTags []orm.Params, count int64) {
	contenTag := new(ContentTag)
	o := orm.NewOrm()
	qs := o.QueryTable(contenTag)
	count, _ = qs.Count()
	qs.Filter("cid", cid).Values(&contentTags)
	return contentTags, count
}
