package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/web_demo/v2/config"
	"github.com/web_demo/v2/log"
	"github.com/web_demo/v2/model/common"
	"net/http"
	"os"
	"path"
	"strings"
)

// StaticCheckMiddleware 获取static前确认是否存在，不存在就返回失败，并触发向数据管理节点获取数据
func StaticCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		middlewareImpl(c)
	}
}

func middlewareImpl(c *gin.Context) {
	pathList := strings.Split(c.Request.URL.Path, "/"+config.Config.GetString("static.relativePath")+"/")
	if len(pathList) < 2 {
		c.Abort()
		log.Sugar.Info("request path error")
		c.JSON(http.StatusBadRequest, common.Response{Msg: "Request path error"})
		return
	}
	fileName := pathList[len(pathList)-1]
	filePath := path.Join(config.Config.GetString("static.root"), fileName)

	// 检查文件是否存在
	// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		c.Abort()
		log.Sugar.Info("not found")
		c.JSON(http.StatusNotFound, common.Response{Msg: "File not exist"})
		// todo 添加文件不存在的处理逻辑
		return
	}
	c.Next()
}
