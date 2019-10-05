package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"gwt/base/httper"
	"gwt/myerr"
	"gwt/service"
	"gwt/val"
)

// AddUser 新增用户
// @Summary 新增用户
// @Description 新增用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body val.AddUserReq true "用户"
// @Success 200 {object} api.BaseRespBody
// @Router /user [post]
func AddUser(ctx *gin.Context) {
	addUserReq := &val.AddUserReq{}
	if httper.BindAndCheck(ctx, addUserReq) {
		return
	}

	err := service.AddUser(addUserReq)
	httper.HandleResponse(ctx, err, nil)
}

// RemoveUser 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body val.DeleteUserReq true "用户"
// @Success 200 {object} api.BaseRespBody
// @Router /user [delete]
func RemoveUser(ctx *gin.Context) {
	deleteUserReq := &val.DeleteUserReq{}
	if httper.BindAndCheck(ctx, deleteUserReq) {
		return
	}

	err := service.RemoveUser(deleteUserReq.Id)
	httper.HandleResponse(ctx, err, nil)
}

// UpdateUser 修改用户
// @Summary 修改用户
// @Description 修改用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body val.UpdateUserReq true "用户"
// @Success 200 {object} api.BaseRespBody
// @Router /user [put]
func UpdateUser(ctx *gin.Context) {
	updateUserReq := &val.UpdateUserReq{}
	if httper.BindAndCheck(ctx, updateUserReq) {
		return
	}

	err := service.UpdateUser(updateUserReq)
	httper.HandleResponse(ctx, err, nil)
}

// GetUser 查询用户 单个
// @Summary 查询用户 单个
// @Description 查询用户 单个
// @Tags User
// @Produce  json
// @Param id path int true "用户id"
// @Success 200 {object} api.GetUserRespAPI
// @Router /user/{id} [get]
func GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		httper.HandleResponse(ctx, myerr.NewParameterError("id为必填参数"), nil)
		return
	}

	user, err := service.GetUser(id)
	httper.HandleResponse(ctx, err, user)
}

// GetUsers 查询用户 列表
// @Summary 查询用户 列表
// @Description 查询用户 列表
// @Tags User
// @Produce  json
// @Param username query string false "用户名"
// @Param nickname query string false "昵称"
// @Param age query string false "年龄"
// @Success 200 {object} api.GetUsersRespAPI
// @Router /users [get]
func GetUsers(ctx *gin.Context) {
	getUsersReq := &val.GetUsersReq{}
	if httper.BindAndCheck(ctx, getUsersReq) {
		return
	}

	users, err := service.GetUsers(getUsersReq)
	httper.HandleResponse(ctx, err, users)
}

// GetUsersPage 查询用户 分页
// @Summary 查询用户 分页
// @Description 查询用户 分页
// @Tags User
// @Produce  json
// @Param page_num query string true "页码"
// @Param page_size query string true "每页大小"
// @Param username query string false "用户名"
// @Param nickname query string false "昵称"
// @Param age query string false "年龄"
// @Success 200 {object} api.GetUsersWithPageAPI
// @Router /users/page [get]
func GetUsersPage(ctx *gin.Context) {
	getUsersWithPageReq := &val.GetUsersWithPageReq{}
	if httper.BindAndCheck(ctx, getUsersWithPageReq) {
		return
	}

	pageModel, err := service.GetUsersPage(getUsersWithPageReq)
	httper.HandleResponse(ctx, err, pageModel)
}
