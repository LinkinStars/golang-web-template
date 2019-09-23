package router

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"controller"
)

// 初始化路由
func InitRouter(port string) {
	r := gin.New()

	// 生产阶段 注释掉下面的gin的Recovery 使用zap输出到固定文件进行格式化并存放堆栈信息
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	// 开发阶段 使用gin的Recovery将日志格式化输出到控制台
	r.Use(gin.Recovery())

	// 将api调用记录全部存放到zap日志中去
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, false))

	// 添加用户
	r.POST("/gwt/api/v1/user", controller.AddUser)
	// 删除用户
	r.DELETE("/gwt/api/v1/user", controller.DeleteUser)
	// 修改用户
	r.PUT("/gwt/api/v1/user", controller.UpdateUser)
	// id查询单用户
	r.GET("/gwt/api/v1/user/:id", controller.GetUser)
	// 多条件分页查询
	r.POST("/gwt/api/v1/users", controller.ListUser)

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
