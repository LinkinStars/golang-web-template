package api

import "gwt/val"

// GetUserRespAPI 查询用户 单个
type GetUserRespAPI struct {
	// 返回码
	Code int `json:"code"`
	// 返回信息
	Message string `json:"message"`
	// 内容
	Data val.GetUserResp `json:"data"`
}

// GetUsersRespAPI 查询用户列表 全部
type GetUsersRespAPI struct {
	// 返回码
	Code int `json:"code"`
	// 返回信息
	Message string `json:"message"`
	// 内容
	Data []val.GetUserResp `json:"data"`
}

// GetUsersWithPageAPI 查询用户列表 分页
type GetUsersWithPageAPI struct {
	// 返回码
	Code int `json:"code"`
	// 返回信息
	Message string `json:"message"`
	// 内容
	Data UsersPageModelAPI `json:"content"`
}

// UsersPageModelAPI 查询用户分页内容
type UsersPageModelAPI struct {
	// 页码
	PageNum int `json:"page_num"`
	// 每页大小
	PageSize int `json:"page_size"`
	// 总页数
	TotalPages int64 `json:"total_pages"`
	// 总记录条数
	TotalRecords int64 `json:"total_records"`
	// 数据
	Records []val.GetUserResp `json:"records"`
}
