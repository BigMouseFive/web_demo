package demo

import (
	"github.com/gin-gonic/gin"
)

// [IMPORTANT] 本文件只是为了在swagger中显示登录/登出两个接口。实际的登录/登出接口参考middleware.jwt.go

type (
	LoginApi struct{}
)

// Login godoc
// @Tags 登录/登出
// @Summary 登录
// @Router /demo/login [post]
// @Description 登录
// @Param user_info body demo.LoginRequest  true "登录信息"
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
func (l *LoginApi) Login(c *gin.Context) {
}

// Logout godoc
// @Tags 登录/登出
// @Summary 登出
// @Router /demo/login [delete]
// @Description 登出
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
func (l *LoginApi) Logout(c *gin.Context) {
}
