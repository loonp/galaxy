package service

import (
	//"github.com/astaxie/beego/orm"
	m "github.com/loonp/galaxy/models"
)

func AddTag(name string) (id int64, err error) {
	id, err = m.AddTag(name)
	return id, err
}

func AddTagWithRead(name string) (id int64, err error) {
	var tag m.Tag
	tag, err = m.GetTagByName(name)
	if err != nil {
		id, err = m.AddTag(name)
	} else {
		id = tag.Id
	}
	return id, err
}

func GetTagById(id int64) (tag m.Tag, err error) {
	tag, err = m.GetTagById(id)
	return tag, err
}

func GetTagByName(name string) (tag m.Tag, err error) {
	tag, err = m.GetTagByName(name)
	return tag, err
}
