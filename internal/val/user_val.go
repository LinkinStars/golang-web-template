package val

import "time"

// AddUserReq 新增用户请求结构
type AddUserReq struct {
	// 用户名
	Username string `validate:"required,gt=0,lte=32" comment:"用户名" json:"username"`
	// 昵称
	Nickname string `validate:"required,gt=0,lte=16" comment:"昵称" json:"nickname"`
	// 年龄
	Age int `validate:"required" comment:"年龄" json:"age"`
}

// RemoveUserReq 删除用户请求结构
type RemoveUserReq struct {
	// 用户ID
	ID int `validate:"required" comment:"用户ID" json:"id"`
}

// UpdateUserReq 修改用户请求结构
type UpdateUserReq struct {
	// 用户ID
	ID int `validate:"required" comment:"用户ID" json:"id"`
	// 用户名
	Username string `validate:"omitempty,gt=0,lte=32" comment:"用户名" json:"username"`
	// 昵称
	Nickname string `validate:"omitempty,gt=0,lte=16" comment:"昵称" json:"nickname"`
	// 年龄
	Age int `validate:"omitempty" comment:"年龄" json:"age"`
}

// GetUsersReq 查询用户列表 全部 请求结构
type GetUsersReq struct {
	// 用户名
	Username string `validate:"omitempty,gt=0,lte=32" comment:"用户名" form:"username"`
	// 昵称
	Nickname string `validate:"omitempty,gt=0,lte=16" comment:"昵称" form:"nickname"`
	// 年龄
	Age int `validate:"omitempty" comment:"年龄" form:"age"`
}

// GetUsersWithPageReq 查询用户列表 分页 请求结构
type GetUsersWithPageReq struct {
	// 页码
	Page int `validate:"required,min=1" comment:"页码" form:"page"`
	// 每页大小
	PageSize int `validate:"required,min=1" comment:"每页大小" form:"page_size"`
	// 用户名
	Username string `validate:"omitempty,gt=0,lte=32" comment:"用户名" form:"username"`
	// 昵称
	Nickname string `validate:"omitempty,gt=0,lte=16" comment:"昵称" form:"nickname"`
	// 年龄
	Age int `validate:"omitempty" comment:"年龄" form:"age"`
}

// GetUserResp 查询用户 返回结构
type GetUserResp struct {
	// 用户ID
	ID int `json:"id"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 修改时间
	UpdatedAt time.Time `json:"updated_at"`
	// 用户名
	Username string `json:"username"`
	// 昵称
	Nickname string `json:"nickname"`
	// 年龄
	Age int `json:"age"`
}
