package controller

import (
	"strconv"

	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/go-scaffold/mistake"
	"github.com/LinkinStars/golang-web-template/internal/base/httper"
	"github.com/LinkinStars/golang-web-template/internal/biz"
	"github.com/LinkinStars/golang-web-template/internal/val"
	"github.com/gin-gonic/gin"
)

// UserController 123
type UserController struct {
	uc  *biz.UserUseCase
	log logger.Logger
}

// NewUserController new a greeter service
func NewUserController(uc *biz.UserUseCase, log logger.Logger) *UserController {
	return &UserController{uc: uc, log: log}
}

// AddUser 新增用户
// @Summary 新增用户
// @Description 新增用户
// @Tags User
// @Accept json
// @Produce json
// @Param data body val.AddUserReq true "用户"
// @Success 200 {object} httper.RespBody
// @Router /user [post]
func (u *UserController) AddUser(ctx *gin.Context) {
	addUserReq := &val.AddUserReq{}
	if httper.BindAndCheck(ctx, addUserReq) {
		return
	}

	err := u.uc.AddUser(ctx, addUserReq)
	httper.HandleResponse(ctx, err, nil)
}

// RemoveUser 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags User
// @Accept json
// @Produce json
// @Param data body val.RemoveUserReq true "用户"
// @Success 200 {object} httper.RespBody
// @Router /user [delete]
func (u *UserController) RemoveUser(ctx *gin.Context) {
	removeUserReq := &val.RemoveUserReq{}
	if httper.BindAndCheck(ctx, removeUserReq) {
		return
	}

	err := u.uc.RemoveUser(ctx, removeUserReq.ID)
	httper.HandleResponse(ctx, err, nil)
}

// UpdateUser 修改用户
// @Summary 修改用户
// @Description 修改用户
// @Tags User
// @Accept json
// @Produce json
// @Param data body val.UpdateUserReq true "用户"
// @Success 200 {object} httper.RespBody
// @Router /user [put]
func (u *UserController) UpdateUser(ctx *gin.Context) {
	updateUserReq := &val.UpdateUserReq{}
	if httper.BindAndCheck(ctx, updateUserReq) {
		return
	}

	err := u.uc.UpdateUser(ctx, updateUserReq)
	httper.HandleResponse(ctx, err, nil)
}

// GetUser 查询用户 单个
// @Summary 查询用户 单个
// @Description 查询用户 单个
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "用户id"
// @Success 200 {object} httper.RespBody{data=val.GetUserResp}
// @Router /user/{id} [get]
func (u *UserController) GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		httper.HandleResponse(ctx, mistake.BadRequest("", "id为必填参数"), nil)
		return
	}

	resp, err := u.uc.GetUser(ctx, id)
	httper.HandleResponse(ctx, err, resp)
}

// GetUsers 查询用户列表 全部
// @Summary 查询用户列表 全部
// @Description 查询用户列表 全部
// @Tags User
// @Produce json
// @Param username query string false "用户名"
// @Param nickname query string false "昵称"
// @Param age query string false "年龄"
// @Success 200 {object} httper.RespBody{data=[]val.GetUserResp}
// @Router /users [get]
func (u *UserController) GetUsers(ctx *gin.Context) {
	getUsersReq := &val.GetUsersReq{}
	if httper.BindAndCheck(ctx, getUsersReq) {
		return
	}

	resp, err := u.uc.GetUsers(ctx, getUsersReq)
	httper.HandleResponse(ctx, err, resp)
}

// GetUsersWithPage 查询用户列表 分页
// @Summary 查询用户列表 分页
// @Description 查询用户列表 分页
// @Tags User
// @Produce json
// @Param page query int false "页码 默认为1"
// @Param page_size query int false "每页大小 默认为10"
// @Param username query string false "用户名"
// @Param nickname query string false "昵称"
// @Param age query string false "年龄"
// @Success 200 {object} httper.RespBody{data=pager.PageModel{records=[]val.GetUserResp}}
// @Router /users/page [get]
func (u *UserController) GetUsersWithPage(ctx *gin.Context) {
	getUsersWithPageReq := &val.GetUsersWithPageReq{}
	if httper.BindAndCheck(ctx, getUsersWithPageReq) {
		return
	}

	resp, err := u.uc.GetUsersWithPage(ctx, getUsersWithPageReq)
	httper.HandleResponse(ctx, err, resp)
}
