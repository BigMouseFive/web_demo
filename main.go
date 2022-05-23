package main

import (
	"context"
	"github.com/web_demo/v2/config"
	"github.com/web_demo/v2/database"
	"github.com/web_demo/v2/log"
	"github.com/web_demo/v2/model/demo"
	"github.com/web_demo/v2/mq"
	"github.com/web_demo/v2/router"
	"time"
)

// @basePath /api
// @title           行云智能SaaS平台接口
// @version         1.0
// @description     负责给SaaS平台前端提供数据
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 配置
	config.Create()

	// 日志
	log.Create()

	// 数据库 redis mongodb
	database.CreateRedis()
	database.CreateMongo()

	// 创建默认的超级用户
	CreateSuperUser()

	// mq
	mq.CreateMqttClient()

	// web
	router.Create()
	router.Run()
}

func CreateSuperUser() {
	var superUser demo.UserModel
	superUser.Id = "demo"
	superUser.Name = "超级管理员"
	superUser.Password = "wwssaadd12345"
	superUser.RoleId = "admin"
	superUser.CreateTime = time.Now().Format("20060102T150405Z")
	_, err := database.Coll("user").InsertOne(context.TODO(), superUser)
	if err != nil {
		log.Sugar.Error(err.Error())
	}
}
