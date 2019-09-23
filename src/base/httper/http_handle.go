package httper

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"base/logger"
	"base/mistake"
	"base/validator"
)

// 绑定参数并验证参数
// true: 确实存在问题，会返回400错误
// false: 不存在问题，验证通过
func BindAndCheck(ctx *gin.Context, data interface{}) bool {
	// 参数映射
	if err := ctx.Bind(data); err != nil {
		HandleResponse(ctx, mistake.NewReqErr("请求参数绑定异常"), nil)
		return true
	}

	// 验证参数
	if err := validator.GlobalValidate.Check(data); err != nil {
		HandleResponse(ctx, mistake.NewReqErr(err.Error()), nil)
		return true
	}

	return false
}

// 统一处理异常，统一处理日志，统一处理返回
func HandleResponse(c *gin.Context, err error, data interface{}) {
	// 如果没有错误，就是正常请求
	if err == nil {
		SendResponseSuccess(c, "操作成功", data)
		return
	}

	// 针对不同的错误类型进行处理
	switch err.(type) {
	// 如果是请求参数异常，返回 400 直接返回err的错误信息
	case *mistake.ReqErr:
		SendResponseFail(c, http.StatusBadRequest, err.Error())

	// 服务异常错误，返回 400 返回err的错误message信息，并打印日志记录
	case *mistake.ServiceErr:
		e := err.(*mistake.ServiceErr)
		if e.Err == nil {
			logger.Error(e.Message, "\n", e.Stack)
		} else {
			logger.Error(e.Message, "\n", e.Err, "\n", e.Stack)
		}
		SendResponseFail(c, http.StatusBadRequest, e.Message)

	// 数据层异常，返回 400 打印错误日志记录，这个错误信息不能返回到前端
	case *mistake.DaoErr:
		e := err.(*mistake.DaoErr)
		logger.Error(e.Err, "\n", e.Stack)
		SendResponseFail(c, http.StatusBadRequest, "数据异常")

	// 未知错误，返回 500，证明此错误没有进行包装，直接被返回
	default:
		logger.Error(err, logger.LogStack(0, 0))
		SendResponseFail(c, http.StatusInternalServerError, "服务器端异常")
	}
	return
}
