package demo

import (
	"context"
	"github.com/web_demo/v2/database"
	"github.com/web_demo/v2/log"
	"github.com/web_demo/v2/model/common"
	"github.com/web_demo/v2/model/demo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type UserService struct{}

//GetUsers
//@author: dylan
//@function: GetUsers
//@description: 获取用户信息
//@return: status int, response interface{}
func (s UserService) GetUsers(data demo.GetUsersReqModel) (int, interface{}) {
	results := make([]demo.UserModel, 0)

	// 分页
	var opt options.FindOptions
	if data.Limit > 0 {
		if data.Page > 0 {
			data.Page -= 1
		}
		opt.SetLimit(data.Limit)
		opt.SetSkip(data.Limit * data.Page)
	}

	// 排序
	if data.Sort != "" {
		if data.Sort[0] == '-' {
			opt.SetSort(bson.M{data.Sort[1:]: -1})
		} else if data.Sort[0] == '+' {
			opt.SetSort(bson.M{data.Sort[1:]: 1})
		}
	}

	// 字段匹配
	and := bson.A{}
	if data.Id != "" {
		and = append(and, bson.M{"id": data.Id})
	}
	if data.CorporationId != "" {
		and = append(and, bson.M{"corporation_id": data.CorporationId})
	}
	if data.RoleId != "" {
		and = append(and, bson.M{"role_id": data.RoleId})
	}
	if data.IgnoreRoleId != "" {
		and = append(and, bson.M{"role_id": bson.M{"$ne": data.IgnoreRoleId}})
	}
	if data.Sex > 0 {
		and = append(and, bson.M{"sex": data.Sex})
	}

	// 模糊搜索 data.Search是regex表达式，在前端完成了封装
	if data.Search != "" {
		and = append(and, bson.M{"$or": bson.A{
			bson.M{"id": bson.M{"$regex": primitive.Regex{Pattern: data.Search, Options: "i"}}},
			bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: data.Search, Options: "i"}}},
		}})
	}
	log.Sugar.Info(data)
	log.Sugar.Info(and)
	filter := bson.M{}
	if len(and) > 0 {
		filter["$and"] = and
	}
	// 查找结果
	cursor, err := database.Coll("user").Find(context.TODO(), filter, &opt)
	if err != nil {
		log.Sugar.Error(err.Error())
		return http.StatusBadRequest, common.Response{Msg: err.Error()}
	}

	// 转换成DeviceModel
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Sugar.Error(err.Error())
		return http.StatusBadRequest, common.Response{Msg: err.Error()}
	}

	return http.StatusOK, common.Response{Result: results}
}

//AddUsers
//@author: dylan
//@function: AddUsers
//@description: 增加用户信息
//@return: status int, response interface{}
func (s UserService) AddUsers(data []demo.UserModel) (int, interface{}) {
	log.Sugar.Info(data)
	newData := make([]interface{}, 0)
	for _, v := range data {
		newData = append(newData, v)
	}
	res, err := database.Coll("user").InsertMany(context.TODO(), newData)
	// InsertMany: 按列表顺序插入，遇到了插入失败的条目就返回error，并且不会继续插入列表中待插入条目
	successIds := make([]string, 0)
	for i := 0; i < len(res.InsertedIDs); i += 1 {
		successIds = append(successIds, data[i].Id)
	}

	if err != nil {
		log.Sugar.Error(err.Error())
		return http.StatusBadRequest, common.Response{Msg: err.Error(), Result: successIds}
	}

	return http.StatusOK, common.Response{Result: successIds}
}

//UpdateUsers
//@author: dylan
//@function: UpdateUsers
//@description: 修改用户信息
//@return: status int, response interface{}
func (s UserService) UpdateUsers(data []common.RequestUpdate) (int, interface{}) {
	// 循环处理
	successIds := make([]string, 0)
	var lastErr string
	for _, value := range data {
		log.Sugar.Info(value)
		_, err := database.Coll("user").UpdateOne(context.TODO(), bson.M{"id": value.Id}, bson.M{"$set": value.Update})
		if err == nil {
			successIds = append(successIds, value.Id)
		} else {
			lastErr = err.Error()
		}
	}

	// 如果修改成功的数量不等于修改用户信息列表的长度则反馈bad_request
	if len(successIds) != len(data) {
		return http.StatusBadRequest, common.Response{Msg: "last error: " + lastErr, Result: successIds}
	}

	return http.StatusOK, common.Response{Result: successIds}
}

//DeleteUsers
//@author: dylan
//@function: DeleteUsers
//@description: 删除用户信息
//@return: status int, response interface{}
func (s UserService) DeleteUsers(data common.RequestIdList) (int, interface{}) {
	// 根据account删除user信息 id in data.IdList
	res, err := database.Coll("user").DeleteMany(context.TODO(),
		bson.M{"id": bson.M{"$in": data.IdList}})
	if err != nil {
		log.Sugar.Error(err.Error())
		return http.StatusBadRequest, common.Response{Msg: err.Error()}
	}
	return http.StatusOK, common.Response{Result: res.DeletedCount}
}

var (
	UserSvr UserService
)
