package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/web_demo/v2/log"
	"github.com/web_demo/v2/model/common"
	"github.com/web_demo/v2/model/demo"
	service "github.com/web_demo/v2/service/demo"
	"net/http"
	"strings"
)

type (
	UserApi struct{}
)

// GetUsers godoc
// @Tags 用户
// @Summary 获取用户信息
// @Router /demo/users [get]
// @Description 获取用户信息
// @Param param query demo.GetUsersReqModel false "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{msg=string,result=[]demo.UserModel}
// @Failure 400 {object} common.Response
// @Security ApiKeyAuth
func (a UserApi) GetUsers(c *gin.Context) {
	var data demo.GetUsersReqModel
	if err := c.ShouldBindQuery(&data); err != nil {
		log.Sugar.Error(err.Error())
		c.JSON(http.StatusBadRequest, common.Response{Msg: err.Error()})
	}

	c.JSON(service.UserSvr.GetUsers(data))
}

// AddUsers godoc
// @Tags 用户
// @Summary 增加用户信息
// @Router /demo/users [put]
// @Description 增加用户信息
// @Param user_list body []demo.UserModel true "新增用户信息列表"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{msg=string,result=[]string} "成功列表"
// @Failure 400 {object} common.Response
// @Security ApiKeyAuth
func (a UserApi) AddUsers(c *gin.Context) {
	var data []demo.UserModel
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Sugar.Error(err.Error())
		c.JSON(http.StatusBadRequest, common.Response{Msg: err.Error()})
		return
	}
	c.JSON(service.UserSvr.AddUsers(data))
}

// UpdateUsers godoc
// @Tags 用户
// @Summary 修改用户信息
// @Router /demo/users [post]
// @Description 修改用户信息
// @Param user_list body []common.RequestUpdate true "修改用户信息列表"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{msg=string,result=[]string} "成功列表"
// @Failure 400 {object} common.Response
// @Security ApiKeyAuth
func (a UserApi) UpdateUsers(c *gin.Context) {
	var data []common.RequestUpdate
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Sugar.Error(err.Error())
		c.JSON(http.StatusBadRequest, common.Response{Msg: err.Error()})
		return
	}
	c.JSON(service.UserSvr.UpdateUsers(data))
}

// DeleteUsers godoc
// @Tags 用户
// @Summary 删除用户信息
// @Router /demo/users [delete]
// @Description 删除用户信息
// @Param id_list query common.RequestIdList true "用户ID列表"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Security ApiKeyAuth
func (a UserApi) DeleteUsers(c *gin.Context) {
	idList, ok := c.GetQuery("id_list")
	if !ok {
		c.JSON(http.StatusBadRequest, common.Response{Msg: "Query 'id_list' not exist."})
		return
	}
	var data common.RequestIdList
	data.IdList = strings.Split(idList, ",")
	c.JSON(service.UserSvr.DeleteUsers(data))
}
