package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"base/httper"
	"service"
	"val"
)

// @Summary 新增用户
// @Description 新增用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body val.AddUserReq true "用户"
// @Success 200 {object} httper.RespBody
// @Router /user [post]
func AddUser(ctx *gin.Context) {
	addUserReq := &val.AddUserReq{}
	if httper.BindAndCheck(ctx, addUserReq) {
		return
	}

	err := service.AddUser(addUserReq)
	httper.HandleResponse(ctx, err, nil)
}

// @Summary 删除用户
// @Description 删除用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body val.DeleteUserReq true "用户"
// @Success 200 {object} httper.RespBody
// @Router /user [delete]
func DeleteUser(ctx *gin.Context) {
	deleteUserReq := &val.DeleteUserReq{}
	if httper.BindAndCheck(ctx, deleteUserReq) {
		return
	}

	err := service.DeleteUser(deleteUserReq.Id)
	httper.HandleResponse(ctx, err, nil)
}

// @Summary 修改用户
// @Description 修改用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body val.UpdateUserReq true "用户"
// @Success 200 {object} httper.RespBody
// @Router /user [put]
func UpdateUser(ctx *gin.Context) {
	updateUserReq := &val.UpdateUserReq{}
	if httper.BindAndCheck(ctx, updateUserReq) {
		return
	}

	err := service.UpdateUser(updateUserReq)
	httper.HandleResponse(ctx, err, nil)
}

// @Summary 根据id查询用户
// @Description 根据id查询用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "用户id"
// @Success 200 {object} val.GetUserResponse
// @Router /user/{id} [get]
func GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		httper.HandleResponse(ctx, errors.New("参数错误"), nil)
		return
	}

	u, err := service.GetUser(id)
	httper.HandleResponse(ctx, err, u)
}

// @Summary 分页查询用户
// @Description 分页查询用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body val.UpdateUserReq true "分页用户条件信息"
// @Success 200 {object} val.GetUsersReq
// @Router /users [get]
func ListUser(ctx *gin.Context) {
	getUsersReq := &val.GetUsersReq{}
	if httper.BindAndCheck(ctx, getUsersReq) {
		return
	}

	pageModel, err := service.ListUser(getUsersReq)
	httper.HandleResponse(ctx, err, pageModel)
}
