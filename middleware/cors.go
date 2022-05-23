package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	CorsMW gin.HandlerFunc
)

// CreateCors 创建cors
func CreateCors() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	CorsMW = cors.New(config)
}
