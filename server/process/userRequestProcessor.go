package process

import (
	"Go_ChatRoom/common/message"
	"Go_ChatRoom/server/model"
	"Go_ChatRoom/server/service"
	"Go_ChatRoom/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserRequestProcessor struct {
	Conn   net.Conn
	UserId int
}

func (this *UserRequestProcessor) LoginProcess(mes *message.Message) (err error) {

	// Unmarshal mes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("userRequestProcessor LoginProcess json.Unmarshal fail err=", err)
		return
	}

	// Declare a response message
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	// Use DAO to handle login request
	user, err := service.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	// Send the response
	// 1. If it is an error, we need to send error message to client
	// 2. Success, then we need to return the status code and user info to client
	if err != nil {

		if err == model.ErrUserNotExist {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ErrInvalidPassword {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}

	} else {

		loginResMes.Code = 200
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)

		this.NotifyOthersOnlineUser(loginMes.UserId)

		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "Login Success!")
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("userRequestProcessor LoginProcess json.Marshal fail", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("userRequestProcessor LoginProcess json.Marshal fail", err)
		return
	}

	tf := &utils.MessageTransformer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}

func (this *UserRequestProcessor) RegisterProcess(mes *message.Message) (err error) {

	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	err = service.MyUserDao.AddUser(&registerMes.User)

	if err != nil {
		if err == model.ErrUserExist {
			registerResMes.Code = 505
			registerResMes.Error = model.ErrUserExist.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "Unknown failure when register..."
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	tf := &utils.MessageTransformer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return

}

func (this *UserRequestProcessor) NotifyMeOnline(userId int) {

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return 
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return 
	}

	tf := &utils.MessageTransformer{
		Conn : this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}

}

func (this *UserRequestProcessor) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}

		up.NotifyMeOnline(userId)
	}
}
