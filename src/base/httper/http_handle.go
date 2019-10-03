package httper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"base/logger"
	"base/validator"
	"myerr"
)

// HandleResponse 统一处理异常，统一处理日志，统一处理返回
func HandleResponse(c *gin.Context, err error, data interface{}) {
	// 如果没有错误，就是正常请求
	if err == nil {
		SendSuccessResp(c, "操作成功", data)
		return
	}

	// 针对不同的错误类型进行处理
	switch errors.Cause(err).(type) {
	case *myerr.ParameterError:
		SendFailResp(c, http.StatusBadRequest, err.Error())
	default:
		logStackInfo(err)
		SendFailResp(c, http.StatusInternalServerError, err.Error())
	}

	return
}

// BindAndCheck 绑定参数并验证参数
// true: 确实存在问题，会返回400错误
// false: 不存在问题，验证通过
func BindAndCheck(ctx *gin.Context, data interface{}) bool {
	// 参数映射
	if err := ctx.Bind(data); err != nil {
		// 当需要的参数类型不匹配的时候会出现错误，属于开发测试时就应该发现的异常
		HandleResponse(ctx, myerr.NewParameterError("绑定参数异常"), nil)
		return true
	}

	// 验证参数
	if err := validator.GlobalValidator.Check(data); err != nil {
		HandleResponse(ctx, myerr.NewParameterError(err.Error()), nil)
		return true
	}

	return false
}

// stackTracer 堆栈信息接口用于内部转换使用
type stackTracer interface {
	StackTrace() errors.StackTrace
}

// logStackInfo 打印错误
func logStackInfo(err error) {
	// 如果包含堆栈信息的封装就打印出相关堆栈信息，如果没有封装就打印原本错误信息
	if e, ok := err.(stackTracer); ok {
		stacks := e.StackTrace()
		var stackEntries []string
		for _, v := range stacks {
			stackEntries = append(stackEntries, fmt.Sprintf("%+v", v))
		}
		logger.Error(stackEntries)
	} else {
		logger.Error(err)
	}
}
