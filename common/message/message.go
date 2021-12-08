package message

const (
	LoginMesType                    = "LoginesType"
	RegisterMesType                 = "RegisteresType"
	LoginResMesType                 = "LoginResesType"
	RegisterResMesType        		= "RegisterResMs"
	NotifyUserStatusMesType 		= "NotifyUserStatusMes"
	SmsMesType          			= "SmsMes"
)

const (
	UserOnline = iota
	UserOffline 
	UserBusyStatus 
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code    int    `json:"code"` // status code : 500 user not registered; 200 success
	Error   string `json:"error"`
	UsersId []int
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"` 
	Status int `json:"status"` 
}

type SmsMes struct {
	Content string `json:"content"` 
	User 
}