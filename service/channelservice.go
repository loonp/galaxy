package service

import (
	"github.com/astaxie/beego/orm"
	m "github.com/loonp/galaxy/models"
)

func GetChannelList(page int64, page_size int64, sort ...string) (channels []orm.Params, count int64) {
	o := orm.NewOrm()
	channel := new(m.Channel)
	qs := o.QueryTable(channel)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort...).Values(&channels)
	count, _ = qs.Count()
	return channels, count
}

func GetAllChannel() (channels []orm.Params, count int64) {
	o := orm.NewOrm()
	channel := new(m.Channel)
	qs := o.QueryTable(channel)
	qs.Values(&channels)
	count, _ = qs.Count()
	return channels, count

}

func AddChannel(channelName string, channelType int8, createId int64) (id int64, err error) {
	id, err = m.AddChannel(channelName, channelType, createId)
	return id, err
}

func GetChannelById(channelId int64) (channel m.Channel, err error) {
	channel, err = m.GetChannelById(channelId)
	return channel, err
}

func UpdateChannel(channelId int64, channelName string, createId int64, channelType int8, channelLevel int32, channelParentid int64) {
	channel := m.Channel{Id: channelId, ChannelName: channelName, CreateId: createId, ChannelType: channelType, ChannelLevel: channelLevel, ChannelParentId: channelParentid}
	m.UpdateChannel(channel)
}

func DelChannel(channelId int64) (num int64, err error) {
	num, err = m.DelChannel(channelId)
	return num, err
}
