package router

import (
	"github.com/LinkinStars/golang-web-template/internal/controller"
	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	uc *controller.UserController
}

// NewUserRoute new a greeter service
func NewUserRoute(uc *controller.UserController) *UserRoute {
	return &UserRoute{uc: uc}
}

// RegisterUserRoute 注册路由
func (u *UserRoute) RegisterUserRoute(r *gin.RouterGroup) {
	r.POST("/user", u.uc.AddUser)
	r.DELETE("/user", u.uc.RemoveUser)
	r.PUT("/user", u.uc.UpdateUser)
	r.GET("/user/:id", u.uc.GetUser)
	r.GET("/users", u.uc.GetUsers)
	r.GET("/users/page", u.uc.GetUsersWithPage)
}
