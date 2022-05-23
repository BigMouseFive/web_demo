package common

type RequestUpdate struct {
	Id     string                 `json:"id" bson:"id"`
	Update map[string]interface{} `json:"update" bson:"update"`
}
