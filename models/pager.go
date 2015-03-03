package models

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Pager struct {
	Page     int64
	Totalnum int64
	Pagesize int64
	urlpath  string
	urlquery string
	nopath   bool
}

func NewPager(page, totalnum, pagesize int64, url string, nopath ...bool) *Pager {
	p := new(Pager)
	p.Page = page
	p.Totalnum = totalnum
	p.Pagesize = pagesize

	arr := strings.Split(url, "?")
	p.urlpath = arr[0]
	if len(arr) > 1 {
		p.urlquery = "?" + arr[1]
	} else {
		p.urlquery = ""
	}

	if len(nopath) > 0 {
		p.nopath = nopath[0]
	} else {
		p.nopath = false
	}

	return p
}

func (this *Pager) url(page int64) string {
	if this.nopath { //不使用目录形式
		if this.urlquery != "" {
			return fmt.Sprintf("%s%s&page=%d", this.urlpath, this.urlquery, page)
		} else {
			return fmt.Sprintf("%s?page=%d", this.urlpath, page)
		}
	} else {
		return fmt.Sprintf("%s/page/%d%s", this.urlpath, page, this.urlquery)
	}
}

func (this *Pager) ToString() string {
	if this.Totalnum <= this.Pagesize {
		return ""
	}

	var buf bytes.Buffer
	var from, to, linknum, offset, totalpage int64

	offset = 5
	linknum = 10

	totalpage = int64(math.Ceil(float64(this.Totalnum) / float64(this.Pagesize)))

	if totalpage < linknum {
		from = 1
		to = totalpage
	} else {
		from = this.Page - offset
		to = from + linknum
		if from < 1 {
			from = 1
			to = from + linknum - 1
		} else if to > totalpage {
			to = totalpage
			from = totalpage - linknum + 1
		}
	}

	buf.WriteString("<span class=\"pageRecordClass\" id=\"pageRecord\">")
	buf.WriteString("[总数:<b>" + fmt.Sprintf("%d", this.Totalnum) + "</b>条</span>]")
	if this.Page == 1 {
		if totalpage == 1 {
			buf.WriteString("[<a href='#'>首页</a>] ")
			buf.WriteString("[<a href='#'>上一页</a>] ")
			buf.WriteString("[<a href='#'>下一页</a>] ")
			buf.WriteString("[<a href='#'>尾页</a>] ")
		} else if totalpage > 1 {
			buf.WriteString("[<a href='#'>首页</a>] ")
			buf.WriteString("[<a href='#'>上一页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", this.Page+1) + ")'>下一页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", totalpage) + ")'>尾页</a>] ")
		}
	} else if this.Page > 1 {
		if totalpage == this.Page {
			buf.WriteString("[<a href='javascript:jump(1)'>首页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", this.Page-1) + ")'>上一页</a>] ")
			buf.WriteString("[<a href='#'>下一页</a>] ")
			buf.WriteString("[<a href='#'>尾页</a>] ")
		} else {
			buf.WriteString("[<a href='javascript:jump(1)'>首页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", this.Page-1) + ")'>上一页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", this.Page+1) + ")'>下一页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", totalpage) + ")'>尾页</a>] ")
		}
	}

	return buf.String()
}

func (this *Pager) ToStringSearch(channelId int64) string {
	if this.Totalnum <= this.Pagesize {
		return ""
	}

	var buf bytes.Buffer
	var from, to, linknum, offset, totalpage int64

	offset = 5
	linknum = 10

	totalpage = int64(math.Ceil(float64(this.Totalnum) / float64(this.Pagesize)))

	if totalpage < linknum {
		from = 1
		to = totalpage
	} else {
		from = this.Page - offset
		to = from + linknum
		if from < 1 {
			from = 1
			to = from + linknum - 1
		} else if to > totalpage {
			to = totalpage
			from = totalpage - linknum + 1
		}
	}

	buf.WriteString("<span class=\"pageRecordClass\" id=\"pageRecord\">")
	buf.WriteString("[总数:<b>" + fmt.Sprintf("%d", this.Totalnum) + "</b>条</span>]")
	if this.Page == 1 {
		if totalpage == 1 {
			buf.WriteString("[<a href='#'>首页</a>] ")
			buf.WriteString("[<a href='#'>上一页</a>] ")
			buf.WriteString("[<a href='#'>下一页</a>] ")
			buf.WriteString("[<a href='#'>尾页</a>] ")
		} else if totalpage > 1 {
			buf.WriteString("[<a href='#'>首页</a>] ")
			buf.WriteString("[<a href='#'>上一页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", this.Page+1) + "," + strconv.Itoa(int(channelId)) + ")'>下一页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", totalpage) + "," + strconv.Itoa(int(channelId)) + ")'>尾页</a>] ")
		}
	} else if this.Page > 1 {
		if totalpage == this.Page {
			buf.WriteString("[<a href='javascript:jump(1)'>首页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", this.Page-1) + "," + strconv.Itoa(int(channelId)) + ")'>上一页</a>] ")
			buf.WriteString("[<a href='#'>下一页</a>] ")
			buf.WriteString("[<a href='#'>尾页</a>] ")
		} else {
			buf.WriteString("[<a href='javascript:jump(1)'>首页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", this.Page-1) + "," + strconv.Itoa(int(channelId)) + ")'>上一页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", this.Page+1) + "," + strconv.Itoa(int(channelId)) + ")'>下一页</a>] ")
			buf.WriteString("[<a href='javascript:jump(" + fmt.Sprintf("%d", totalpage) + "," + strconv.Itoa(int(channelId)) + ")'>尾页</a>] ")
		}
	}

	return buf.String()
}

func (this *Pager) ToStringIndex(channelId int64) string {
	if this.Totalnum <= this.Pagesize {
		return ""
	}

	var buf bytes.Buffer
	var from, to, linknum, offset, totalpage int64

	offset = 5
	linknum = 10

	totalpage = int64(math.Ceil(float64(this.Totalnum) / float64(this.Pagesize)))

	if totalpage < linknum {
		from = 1
		to = totalpage
	} else {
		from = this.Page - offset
		to = from + linknum
		if from < 1 {
			from = 1
			to = from + linknum - 1
		} else if to > totalpage {
			to = totalpage
			from = totalpage - linknum + 1
		}
	}

	buf.WriteString("<span class=\"pageRecordClass\" id=\"pageRecord\">")
	buf.WriteString("[总数:<b>" + fmt.Sprintf("%d", this.Totalnum) + "</b>条</span>]")
	if this.Page == 1 {
		if totalpage == 1 {
			buf.WriteString("[<a href='#'>首页</a>] ")
			buf.WriteString("[<a href='#'>上一页</a>] ")
			buf.WriteString("[<a href='#'>下一页</a>] ")
			buf.WriteString("[<a href='#'>尾页</a>] ")
		} else if totalpage > 1 {
			buf.WriteString("[<a href='#'>首页</a>] ")
			buf.WriteString("[<a href='#'>上一页</a>] ")
			buf.WriteString("[<a href='/?page=" + fmt.Sprintf("%d", this.Page+1) + "&channelId=" + strconv.Itoa(int(channelId)) + "'>下一页</a>] ")
			buf.WriteString("[<a href='/?page=" + fmt.Sprintf("%d", totalpage) + "&channelId=" + strconv.Itoa(int(channelId)) + "'>尾页</a>] ")
		}
	} else if this.Page > 1 {
		if totalpage == this.Page {
			buf.WriteString("[<a href='/?page=1'>首页</a>] ")
			buf.WriteString("[<a href='/?page=" + fmt.Sprintf("%d", this.Page-1) + "&channelId=" + strconv.Itoa(int(channelId)) + "'>上一页</a>] ")
			buf.WriteString("[<a href='#'>下一页</a>] ")
			buf.WriteString("[<a href='#'>尾页</a>] ")
		} else {
			buf.WriteString("[<a href='/?page=1)'>首页</a>] ")
			buf.WriteString("[<a href='/?page=" + fmt.Sprintf("%d", this.Page-1) + "&channelId=" + strconv.Itoa(int(channelId)) + "'>上一页</a>] ")
			buf.WriteString("[<a href='/?page=" + fmt.Sprintf("%d", this.Page+1) + "&channelId=" + strconv.Itoa(int(channelId)) + "'>下一页</a>] ")
			buf.WriteString("[<a href='/?page=" + fmt.Sprintf("%d", totalpage) + "&channelId=" + strconv.Itoa(int(channelId)) + "'>尾页</a>] ")
		}
	}

	return buf.String()
}
