package process

import(
	"fmt"
	"net"
	"Go_ChatRoom/common/message"
	"Go_ChatRoom/utils"
	"encoding/json"
)

type UserProcess struct {

}

func(this *UserProcess) Login(userId int, userPwd string) (err error) {

	//Build net connction
	//Return if get errors
	//Defer close the connection
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("Dial error=", err) 
		return
	}
	defer conn.Close()

	//Serialize data
	loginMes := &message.LoginMes {
		UserId : userId,
		UserPwd : userPwd,
	}

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return 
	}

	mes := &message.Message {
		Type : message.LoginMesType,
		Data : string(data),
	}

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return 
	}

	//Send data
	msgTransformer := &utils.MessageTransformer {
		Conn : conn,
	}

	err = msgTransformer.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}

	return 
}