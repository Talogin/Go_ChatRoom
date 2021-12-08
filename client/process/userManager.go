package process

import (
	"Go_ChatRoom/client/model"
	"Go_ChatRoom/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

func outputOnlineUser() {
	//遍历一把 onlineUsers
	fmt.Println("Current online users are listed as below:")
	for id, _ := range onlineUsers {
		//如果不显示自己.
		fmt.Println("user id:\t", id)
	}
}

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
