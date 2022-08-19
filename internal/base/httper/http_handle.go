package httper

import (
	"errors"
	"net/http"

	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/go-scaffold/mistake"
	"github.com/LinkinStars/golang-util/gu"
	"github.com/LinkinStars/golang-web-template/internal/base/validator"
	"github.com/gin-gonic/gin"
)

// HandleResponse 统一处理异常，统一处理日志，统一处理返回
func HandleResponse(c *gin.Context, err error, data interface{}) {
	// 如果没有错误，就是正常请求
	if err == nil {
		c.JSON(http.StatusOK, NewRespBodyData(http.StatusOK, Success, data))
		return
	}

	var myErr *mistake.Error
	// 未知错误
	if !errors.As(err, &myErr) {
		logger.Error(err, "\n", mistake.LogStack(2, 5))
		c.JSON(http.StatusInternalServerError, NewRespBody(http.StatusInternalServerError, UnexpectedError))
		return
	}

	// 服务异常打印日志
	if mistake.IsInternalServer(myErr) {
		logger.Error(myErr)
	}

	c.JSON(myErr.Code, NewRespBodyFromError(myErr))
	return
}

// BindAndCheck 绑定参数并验证参数
// true: 确实存在问题，会返回400错误
// false: 不存在问题，验证通过
func BindAndCheck(ctx *gin.Context, data interface{}) bool {
	// 参数映射
	if err := ctx.ShouldBind(data); err != nil {
		// 当入参格式不正确的会出现，比如int传递为string
		logger.Errorf("http_handle BindAndCheck fail, %s", err.Error())
		HandleResponse(ctx, mistake.New(http.StatusBadRequest, "RequestFormatError", RequestFormatError), nil)
		return true
	}

	// 去除结构体内部数据前后的空格
	gu.TrimStruct(data)

	// 验证参数
	if err := validator.GlobalValidator.Check(data); err != nil {
		HandleResponse(ctx, mistake.New(http.StatusBadRequest, "BadRequest", err.Error()), nil)
		return true
	}
	return false
}
