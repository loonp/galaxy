package service

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/loonp/galaxy/lib"
	m "github.com/loonp/galaxy/models"
	"time"
)

func CheckLogin(username string, password string) (user m.User, err error) {
	user = m.GetUserByUsername(username)
	if user.Id == 0 {
		return user, errors.New("用户不存在")
	}
	if user.Password != lib.Pwdhash(password) {
		return user, errors.New("密码错误")
	}
	return user, nil
}

func AddUser(username string, password string, email string) (id int64, err error) {
	user := m.User{Username: username, Password: lib.Pwdhash(password), Email: email, RegisterTime: time.Now()}
	fmt.Println(user)
	id, err = m.AddUser(user)
	return id, err
}

func GetUserById(id int64) (user m.User) {
	user = m.GetUserById(id)
	return user
}

func UpdateUser(id int64, username string, email string) (err error) {
	user := m.User{Id: id, Username: username, Email: email}
	fmt.Println(user)
	_, err = m.UpdateUser(user)
	return err
}

func GetUserList(page int64, page_size int64, sort ...string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(m.User)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort...).Values(&users)
	count, _ = qs.Count()
	return users, count
}
func DelUser(userId int64) (num int64, err error) {
	num, err = m.DelUser(userId)
	return num, err
}
