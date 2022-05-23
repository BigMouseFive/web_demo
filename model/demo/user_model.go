package demo

import "github.com/web_demo/v2/model/common"

type UserModel struct {
	Id            string      `bson:"id" json:"id"`
	Name          string      `bson:"name" json:"name"`
	Password      string      `bson:"password" json:"password"`
	CreateTime    string      `bson:"create_time" json:"create_time"`
	CorporationId string      `bson:"corporation_id" json:"corporation_id"`
	Phone         string      `bson:"phone" json:"phone"`
	Sex           int         `bson:"sex" json:"sex"` // 0:未知  1:男  2:女
	Email         string      `bson:"email" json:"email"`
	RoleId        string      `bson:"role_id" json:"role_id"`
	Projects      []string    `bson:"projects" json:"projects"`
	Auths         []string    `bson:"auths" json:"auths"`
	Remark        interface{} `bson:"remark" json:"remark"`
}

type GetUsersReqModel struct {
	common.GetCommonModel
	CorporationId string `bson:"corporation_id" json:"corporation_id" form:"corporation_id"` // 绑定的客户编号
	RoleId        string `bson:"role_id" json:"role_id" form:"role_id"`                      // 绑定的角色编号
	Sex           int64  `bson:"sex" json:"sex" form:"sex"`                                  // 性别
	IgnoreRoleId  string `bson:"ignore_role_id" json:"ignore_role_id" form:"ignore_role_id"`
}
