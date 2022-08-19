package server

import (
	"github.com/LinkinStars/go-scaffold/contrib/middleware/gin/recovery"
	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/golang-web-template/internal/base/conf"
	"github.com/LinkinStars/golang-web-template/internal/router"
	"github.com/gin-gonic/gin"
)

// NewHTTPServer new an HTTP s.
func NewHTTPServer(c *conf.Server, log logger.Logger, userRouter *router.UserRoute) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(recovery.Recovery())

	// 需要鉴权的路由
	v1 := r.Group("/api/gwt/v1")

	userRouter.RegisterUserRoute(v1)
	return r
}
