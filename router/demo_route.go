package router

import (
	"github.com/gin-gonic/gin"
	"github.com/web_demo/v2/api/demo"
	"github.com/web_demo/v2/middleware"
)

// InitLoginRoute 初始化登录的路由
func InitLoginRoute(Router *gin.RouterGroup) {
	r := Router.Group("login")
	r.POST("", middleware.JwtMW.LoginHandler)
	r.DELETE("", middleware.JwtMW.LogoutHandler)
}

// InitUserRoute 初始化用户的路由
func InitUserRoute(Router *gin.RouterGroup) {
	r := Router.Group("users")
	a := demo.UserApi{}
	r.GET("", a.GetUsers)
	r.PUT("", a.AddUsers)
	r.POST("", a.UpdateUsers)
	r.DELETE("", a.DeleteUsers)
}
