package demo

import (
	"context"
	"github.com/web_demo/v2/database"
	"github.com/web_demo/v2/log"
	"github.com/web_demo/v2/model/common"
	"github.com/web_demo/v2/model/demo"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type LoginService struct{}

//Login
//@author: dylan
//@function: Login
//@description: 登录
//@return: status int, response interface{}
func (s LoginService) Login(data demo.LoginRequest) (int, interface{}) {
	// 根据account获取user信息 id == data.account
	var result bson.M
	err := database.Coll("user").FindOne(context.TODO(),
		bson.M{"id": data.Account}).Decode(&result)
	if err != nil {
		log.Sugar.Error(err.Error())
		return http.StatusBadRequest, common.Response{Msg: "Account is not exist."}
	}

	// 比较密码是否一致
	if result["password"] != data.Password {
		return http.StatusBadRequest, common.Response{Msg: "Password is not correct."}
	}

	return http.StatusOK, common.Response{}
}

var (
	LoginSvc LoginService
)
