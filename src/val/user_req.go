package val

import (
	"errors"
	"strings"
)

// AddUserReq 添加用户请求
type AddUserReq struct {
	// 用户名
	Username string `validate:"required,gt=6,lt=16" comment:"用户名" json:"username"`
	// 昵称
	Nickname string `validate:"omitempty,gt=6,lt=16" comment:"昵称" json:"nickname"`
	// 年龄
	Age int `validate:"required,min=5,max=99" comment:"年龄" json:"age"`
}

func (u *AddUserReq) Check() error {
	if strings.Contains(u.Username, "@") {
		return errors.New("用户名不能包含@特殊字符")
	}
	return nil
}

// DeleteUserReq 删除用户请求
type DeleteUserReq struct {
	// 用户id
	Id int `validate:"required,min=1" comment:"Id"`
}

// UpdateUserReq 更新用户请求
type UpdateUserReq struct {
	// 用户id
	Id int `validate:"required,min=1" comment:"Id"`
	// 用户名
	Username string `validate:"omitempty,gt=6,lt=16" comment:"用户名" json:"username"`
	// 昵称
	Nickname string `validate:"omitempty,gt=6,lt=16" comment:"昵称" json:"nickname"`
	// 年龄
	Age int `validate:"omitempty,min=5,max=99" comment:"年龄" json:"age"`
}

// GetUsersReq 查询用户列表请求
type GetUsersReq struct {
	// 用户名
	Username string `validate:"omitempty,gt=6,lt=16" comment:"用户名" form:"username"`
	// 昵称
	Nickname string `validate:"omitempty,gt=6,lt=16" comment:"昵称" form:"nickname"`
	// 年龄
	Age int `validate:"omitempty,min=5,max=99" comment:"年龄" form:"age"`
}

// GetUsersWithPageReq 查询用户分页请求
type GetUsersWithPageReq struct {
	// 页码
	PageNum int `validate:"required,min=1" comment:"页码" form:"page_num"`
	// 每页大小
	PageSize int `validate:"required,min=1" comment:"每页大小" form:"page_size"`
	// 用户名
	Username string `validate:"omitempty,gt=6,lt=16" comment:"用户名" form:"username"`
	// 昵称
	Nickname string `validate:"omitempty,gt=6,lt=16" comment:"昵称" form:"nickname"`
	// 年龄
	Age int `validate:"omitempty,min=5,max=99" comment:"年龄" form:"age"`
}

// GetUserResp 查询用户 返回结构
type GetUserResp struct {
	// 用户id
	ID int `json:"id"`
	// 用户名
	Username string `json:"username"`
	// 昵称
	Nickname string `json:"nickname"`
	// 年龄
	Age int `json:"age"`
}
