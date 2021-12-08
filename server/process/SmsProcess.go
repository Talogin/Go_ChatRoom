package process

import (
	"Go_ChatRoom/common/message"
	"Go_ChatRoom/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//这里，还需要过滤到自己,即不要再发给自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}

}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	//创建一个Transfer 实例，发送data
	tf := &utils.MessageTransformer{
		Conn: conn, //
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("send message err=", err)
	}
}
