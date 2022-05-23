package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	gindump "github.com/tpkeeper/gin-dump"
	"github.com/web_demo/v2/config"
	"github.com/web_demo/v2/docs"
	"github.com/web_demo/v2/log"
	"github.com/web_demo/v2/middleware"
)

var (
	Web *gin.Engine
)

// Create 创建WEB服务
func Create() {
	Web = gin.New()

	// 初始化中间件
	middleware.CreateJWT()
	middleware.CreateCors()

	// 日志中间件和异常处理中间件
	Web.Use(gin.Logger()).Use(gin.Recovery())
	// 转储请求和响应的标头/正文的中间件 https://github.com/tpkeeper/gin-dump
	Web.Use(gindump.Dump())
	// 跨域中间件 https://github.com/gin-contrib/cors 增加jwt所需要的header:Authorization
	Web.Use(middleware.CorsMW)

	allApi := Web.Group("api")
	public := allApi.Group("demo")
	InitLoginRoute(public)
	private := allApi.Group("demo")
	// 使用jwt中间件https://github.com/appleboy/gin-jwt
	private.Use(middleware.JwtMW.MiddlewareFunc())
	InitUserRoute(private)

	// 静态资源
	saasStatic := public.Group(config.Config.GetString("static.relativePath"))
	saasStatic.Use(middleware.StaticCheckMiddleware())
	saasStatic.Static("", config.Config.GetString("static.root"))

	// swagger route
	docs.SwaggerInfo.BasePath = "/api"
	Web.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// Run 启动WEB服务
func Run() {
	log.Sugar.Infow("start web ")
	log.Sugar.Fatal(Web.Run(":8888"))
}
