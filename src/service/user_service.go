package service

import (
	"github.com/jinzhu/copier"

	"dao"
	"model"
	"val"
)

// 新增用户
func AddUser(addUserReq *val.AddUserReq) (err error) {
	// 转换
	u := &model.User{}
	if err := copier.Copy(u, addUserReq); err != nil {
		return err
	}

	if err = dao.AddUser(u); err != nil {
		return err
	}
	return nil
}

// 删除用户
func DeleteUser(id int) (err error) {
	if err = dao.DeleteUser(id); err != nil {
		return err
	}
	return nil
}

// 修改用户
func UpdateUser(updateUserReq *val.UpdateUserReq) (err error) {
	// 转换
	u := &model.User{}
	if err := copier.Copy(u, updateUserReq); err != nil {
		return err
	}

	if err = dao.UpdateUser(u); err != nil {
		return err
	}
	return nil
}

// 查询用户
func GetUser(id int) (ur *val.GetUserResponse, err error) {
	u, err := dao.GetUser(id)
	if err != nil {
		return nil, err
	}

	// 转换
	ur = &val.GetUserResponse{}
	if err := copier.Copy(ur, u); err != nil {
		return nil, err
	}
	return ur, nil
}
