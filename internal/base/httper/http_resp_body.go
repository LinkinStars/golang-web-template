package httper

import "github.com/LinkinStars/go-scaffold/mistake"

// RespBody 响应结构体
type RespBody struct {
	// 状态码
	Code int `json:"code"`
	// 错误原因
	Reason string `json:"reason"`
	// 描述
	Message string `json:"message"`
	// 内容
	Data interface{} `json:"data"`
}

// NewRespBody 创建返回数据
func NewRespBody(code int, msg string) *RespBody {
	return &RespBody{
		Code:    code,
		Message: msg,
	}
}

// NewRespBodyFromError 创建返回数据
func NewRespBodyFromError(e *mistake.Error) *RespBody {
	return &RespBody{
		Code:    e.Code,
		Reason:  e.Reason,
		Message: e.Message,
	}
}

// NewRespBodyData 创建返回数据
func NewRespBodyData(code int, msg string, data interface{}) *RespBody {
	return &RespBody{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
