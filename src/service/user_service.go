package service

import (
	"github.com/jinzhu/copier"

	"dao"
	"model"
	"util/pager"
	"val"
)

// AddUser 新增用户
func AddUser(addUserReq *val.AddUserReq) (err error) {
	user := &model.User{}
	_ = copier.Copy(user, addUserReq)
	return dao.AddUser(user)
}

// RemoveUser 删除用户
func RemoveUser(id int) (err error) {
	return dao.RemoveUser(id)
}

// UpdateUser 修改用户
func UpdateUser(updateUserReq *val.UpdateUserReq) (err error) {
	user := &model.User{}
	_ = copier.Copy(user, updateUserReq)
	return dao.AddUser(user)
}

// GetUser 查询用户 单个
func GetUser(id int) (getUserResp *val.GetUserResp, err error) {
	user, err := dao.GetUser(id)
	if err != nil || user == nil {
		return
	}

	getUserResp = &val.GetUserResp{}
	_ = copier.Copy(getUserResp, user)
	return getUserResp, nil
}

// GetUsers 查询用户 列表
func GetUsers(getUsersReq *val.GetUsersReq) (usersResp *[]val.GetUserResp, err error) {
	user := &model.User{}
	_ = copier.Copy(user, getUsersReq)

	users, err := dao.GetUsers(user)
	if err != nil {
		return
	}

	usersResp = &[]val.GetUserResp{}
	_ = copier.Copy(usersResp, users)
	return
}

// GetUsers 查询用户 分页
func GetUsersPage(getUsersWithPageReq *val.GetUsersWithPageReq) (pageModel *pager.PageModel, err error) {
	// 将请求参数转换为查询参数
	user := &model.User{}
	_ = copier.Copy(user, getUsersWithPageReq)
	pageNum := getUsersWithPageReq.PageNum
	pageSize := getUsersWithPageReq.PageSize

	// 分页查询
	users, total, err := dao.GetUsersPage(pageNum, pageSize, user)
	if err != nil {
		return
	}

	// 将查询结果封装返回
	usersResp := &[]val.GetUserResp{}
	_ = copier.Copy(usersResp, users)
	return pager.NewPageModel(pageNum, pageSize, total, usersResp), nil
}
