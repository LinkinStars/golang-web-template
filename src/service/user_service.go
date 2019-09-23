package service

import (
	"github.com/jinzhu/copier"

	"base/logger"
	"base/mistake"
	"base/pager"
	"dao"
	"model"
	"val"
)

// 新增用户
func AddUser(addUserReq *val.AddUserReq) (err error) {
	u := &model.User{}
	_ = copier.Copy(u, addUserReq)

	return dao.AddUser(u)
}

// 删除用户
func DeleteUser(id int) (err error) {
	return dao.DeleteUser(id)
}

// 修改用户
func UpdateUser(updateUserReq *val.UpdateUserReq) (err error) {
	if ur, err := GetUser(updateUserReq.Id); ur == nil || err != nil {
		logger.Warn("修改信息不存在：", updateUserReq.Id)
		return mistake.NewServiceErr(nil, "修改信息不存在")
	}

	u := &model.User{}
	_ = copier.Copy(u, updateUserReq)

	return dao.UpdateUser(u)
}

// 查询用户
func GetUser(id int) (ur *val.GetUserResponse, err error) {
	u, err := dao.GetUser(id)
	if err != nil {
		return nil, err
	}

	// 未查询到信息
	if u.Id == 0 {
		return nil, nil
	}

	ur = &val.GetUserResponse{}
	_ = copier.Copy(ur, u)
	return ur, nil
}

// 分页查询用户
func ListUser(req *val.GetUsersReq) (pageModel *pager.PageModel, err error) {
	return dao.ListUser(req.PageNum, req.PageSize, req.Username, req.Age)
}
