package api

// BaseRespBody 基础返回结构
type BaseRespBody struct {
	// 返回码
	Code int `json:"code"`
	// 返回信息
	Message string `json:"message"`
	// 数据
	Data interface{} `json:"data,omitempty"`
}
