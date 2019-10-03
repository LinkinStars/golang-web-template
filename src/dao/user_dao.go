package dao

import (
	"github.com/pkg/errors"

	"base/db"
	"model"
	"util/pager"
)

// AddUser 添加用户
func AddUser(u *model.User) (err error) {
	_, err = db.Engine.Insert(u)
	return errors.WithStack(err)
}

// RemoveUser 删除用户
func RemoveUser(id int) (err error) {
	_, err = db.Engine.Id(id).Delete(model.User{})
	return errors.WithStack(err)
}

// UpdateUser 修改用户
func UpdateUser(user *model.User) (err error) {
	_, err = db.Engine.Id(user.ID).Update(user)
	return errors.WithStack(err)
}

// GetUser 查询用户 单个
func GetUser(id int) (user *model.User, err error) {
	user = &model.User{}
	exist, err := db.Engine.ID(id).Get(user)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !exist {
		user = nil
	}
	return
}

// GetUsers 查询用户 列表
func GetUsers(user *model.User) (users *[]model.User, err error) {
	users = &[]model.User{}
	err = db.Engine.Find(users, user)
	err = errors.WithStack(err)
	return
}

// GetUsersPage 查询用户 分页
func GetUsersPage(pageNum, pageSize int, user *model.User) (users *[]model.User, total int64, err error) {
	users = &[]model.User{}
	total, err = pager.Help(pageNum, pageSize, users, user, db.Engine.NewSession())
	err = errors.WithStack(err)
	return
}
