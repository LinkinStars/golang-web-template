package val

import (
	"errors"
	"strings"
)

// AddUserReq 添加用户请求
type AddUserReq struct {
	Username string `validate:"required,gt=6,lt=16" comment:"用户名" json:"username"`
	Nickname string `validate:"omitempty,gt=6,lt=16" comment:"昵称" json:"nickname"`
	Age      int    `validate:"required,min=5,max=99" comment:"年龄" json:"age"`
}

func (u *AddUserReq) Check() error {
	if strings.Contains(u.Username, "@") {
		return errors.New("用户名不能包含@特殊字符")
	}
	return nil
}

// DeleteUserReq 删除用户请求
type DeleteUserReq struct {
	Id int `validate:"required,min=1" comment:"Id"`
}

// UpdateUserReq 更新用户请求
type UpdateUserReq struct {
	Id       int    `validate:"required,min=1" comment:"Id"`
	Username string `validate:"omitempty,gt=6,lt=16" comment:"用户名" json:"username"`
	Nickname string `validate:"omitempty,gt=6,lt=16" comment:"昵称" json:"nickname"`
	Age      int    `validate:"omitempty,min=5,max=99" comment:"年龄" json:"age"`
}

// GetUsersReq 查询用户列表请求
type GetUsersReq struct {
	Username string `validate:"omitempty,gt=6,lt=16" comment:"用户名" form:"username"`
	Nickname string `validate:"omitempty,gt=6,lt=16" comment:"昵称" form:"nickname"`
	Age      int    `validate:"omitempty,min=5,max=99" comment:"年龄" form:"age"`
}

// GetUsersWithPageReq 查询用户分页请求
type GetUsersWithPageReq struct {
	// 页码
	PageNum int `validate:"required,min=1" comment:"页码" form:"pageNum"`
	// 每页大小
	PageSize int    `validate:"required,min=1" comment:"每页大小" form:"pageSize"`
	Username string `validate:"omitempty,gt=6,lt=16" comment:"用户名" form:"username"`
	Nickname string `validate:"omitempty,gt=6,lt=16" comment:"昵称" form:"nickname"`
	Age      int    `validate:"omitempty,min=5,max=99" comment:"年龄" form:"age"`
}

// GetUserResp 查询用户 返回结构
type GetUserResp struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
}
