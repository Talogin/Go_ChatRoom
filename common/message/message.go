package message

const (
	LoginMesType = "LoginMesType"
	RegisterMesType = "RegisterMesType"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId int `json:"userId"` 
	UserPwd string `json:"userPwd"` 
	UserName string `json:"userName"` 
}