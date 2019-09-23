package dao

import (
	"github.com/pkg/errors"

	"base/db"
	"base/mistake"
	"base/pager"
	"model"
)

// 添加用户
func AddUser(u *model.User) (err error) {
	_, err = db.DataEngine.Insert(u)
	if err != nil {
		return mistake.NewDaoErr(err)
	}
	return nil
}

// 删除用户
func DeleteUser(id int) (err error) {
	_, err = db.DataEngine.Delete(model.User{})
	if err != nil {
		return mistake.NewDaoErr(err)
	}
	return nil
}

// 修改用户
func UpdateUser(u *model.User) (err error) {
	_, err = db.DataEngine.Id(u.Id).Update(u)
	if err != nil {
		return mistake.NewDaoErr(err)
	}
	return err
}

// 查询用户
func GetUser(id int) (u *model.User, err error) {
	u = &model.User{}
	_, err = db.DataEngine.ID(id).Get(u)
	if err != nil {
		return nil, mistake.NewDaoErr(err)
	}
	return u, nil
}

// 根据查询条件分页用户
func ListUser(page, pageSize int, username string, age int) (pageModel *pager.PageModel, err error) {
	// 设置查询条件
	session := db.DataEngine.Where("username LIKE ?", username+"%").And("age > ?", age)

	// 进行分页查询
	users := &[]model.User{}
	total, err := pager.Help(page, pageSize, users, session)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 计算分页信息并返回
	return pager.NewPageModel(page, pageSize, total, *users), nil
}
