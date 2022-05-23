package common

type GetCommonModel struct {
	Id     string `bson:"id" json:"id" form:"id"`             // 根据id获取
	Page   int64  `bson:"page" json:"page" form:"page"`       // 分页：页码
	Limit  int64  `bson:"limit" json:"limit" form:"limit"`    // 分页：每页显示数量
	Search string `bson:"search" json:"search" form:"search"` // 全文检索，针对id和name
	Sort   string `bson:"sort" json:"sort" form:"sort"`       // 排序 +key/-key
}
