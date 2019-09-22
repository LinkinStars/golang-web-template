package httper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendResponseFail(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, Response{
		Code:    -1,
		Message: message,
	})
}

func SendResponseSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    1,
		Message: message,
		Data:    data,
	})
}
