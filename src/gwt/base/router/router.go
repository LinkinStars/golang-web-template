package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"gwt/controller"
)

// InitRouter 初始化路由
func InitRouter(port string) {
	r := gin.New()

	// 生产阶段 注释掉下面的gin的Recovery 使用zap输出到固定文件进行格式化并存放堆栈信息
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	// 开发阶段 使用gin的Recovery将日志格式化输出到控制台
	r.Use(gin.Recovery())

	// 将api调用记录全部存放到zap日志中去
	r.Use(ginzap.Ginzap(zap.L(), "2006-01-02 15:04:05.000", false))

	r.POST("/gwt/api/v1/user", controller.AddUser)
	r.DELETE("/gwt/api/v1/user", controller.RemoveUser)
	r.PUT("/gwt/api/v1/user", controller.UpdateUser)
	r.GET("/gwt/api/v1/user/:id", controller.GetUser)
	r.GET("/gwt/api/v1/users", controller.GetUsers)
	r.GET("/gwt/api/v1/users/page", controller.GetUsersPage)

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
