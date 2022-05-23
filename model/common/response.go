package common

type Response struct {
	Msg    string      `json:"msg" bson:"msg"`
	Result interface{} `json:"result" bson:"result"`
}
