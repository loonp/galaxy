package service

import (
	"github.com/astaxie/beego/orm"
	m "github.com/loonp/galaxy/models"
	"time"
)

func GetContentList(page int64, page_size int64, sorts ...string) (contents []orm.Params, count int64) {
	o := orm.NewOrm()
	content := new(m.Content)
	qs := o.QueryTable(content)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sorts...).Values(&contents)
	count, _ = qs.Count()
	return contents, count
}

func GetContentListByChannelId(page int64, page_size int64, channelId int64, sorts ...string) (contents []orm.Params, count int64) {
	o := orm.NewOrm()
	content := new(m.Content)
	qs := o.QueryTable(content)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Filter("content_channel_id", channelId).Limit(page_size, offset).OrderBy(sorts...).Values(&contents)
	count, _ = qs.Filter("content_channel_id", channelId).Count()
	return contents, count
}

func AddContent(contentChannelId int64, contentTitle, contentSummary, contentMain string, createId int64, mediaPath string) (id int64, err error) {
	content := m.Content{ContentChannelId: contentChannelId, ContentTitle: contentTitle, ContentSummary: contentSummary, ContentMain: contentMain, CreateId: createId}
	content.CreateTime = time.Now()
	content.UpdateTime = time.Now()
	content.MediaPath = mediaPath
	id, err = m.AddContent(content)
	return id, err
}

func UpdateContent(id, createId, updateId int64, contentChannelId int64, contentTitle, contentSummary, contentMain string, mediaPath string) (err error) {
	//content := m.Content{Id: id, CreateId: createId, UpdateId: updateId, ContentChannelId: contentChannelId, ContentTitle: contentTitle, ContentSummary: contentSummary, ContentMain: contentMain}
	content := m.GetContentById(id)
	content.UpdateId = updateId
	content.ContentChannelId = contentChannelId
	content.ContentTitle = contentTitle
	content.ContentSummary = contentSummary
	content.ContentMain = contentMain
	content.UpdateTime = time.Now()
	content.MediaPath = mediaPath
	_, err = m.UpdateContent(content)
	return err
}

func GetContentById(id int64) (content m.Content) {
	content = m.GetContentById(id)
	return content
}

func DelContent(contentId int64) (num int64, err error) {
	num, err = m.DelContent(contentId)
	return num, err
}
