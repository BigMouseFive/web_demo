package demo

type LoginRequest struct {
	Account  string `json:"account" bson:"account"`
	Password string `json:"password" bson:"password"`
}
