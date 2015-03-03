package service

import (
	"github.com/astaxie/beego/orm"
	m "github.com/loonp/galaxy/models"
)

func AddContentTag(tid int64, cid int64) (id int64, err error) {
	id, err = m.AddContentTag(tid, cid)
	return id, err
}

func GetContentTagsByCid(cid int64) (contentTags []orm.Params, count int64) {
	contentTags, count = m.GetContentTagsByCid(cid)
	return contentTags, count
}
