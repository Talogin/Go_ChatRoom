package main

import (
	"Go_ChatRoom/common/message"
	"Go_ChatRoom/server/process"
	"Go_ChatRoom/utils"
	"fmt"
	"io"
	"net"
)

type UserConnectionHandler struct {
	Conn net.Conn
}

func (this *UserConnectionHandler) serverProcessMes(mes *message.Message) (err error) {

	//mes contains 'type' & 'data'
	switch mes.Type {
		case message.LoginMesType :
			userRP := &process.UserRequestProcessor{
				Conn: this.Conn,
			}
			err = userRP.LoginProcess(mes)
		case message.RegisterMesType :
			userRP := &process.UserRequestProcessor{
				Conn: this.Conn,
			}
			err = userRP.RegisterProcess(mes)
		case message.SmsMesType :
			
		default :
			fmt.Println("The message type does not exist....")
	}
	return
}

func (this *UserConnectionHandler) process() (err error) {

	msgTransformer := &utils.MessageTransformer{
		Conn: this.Conn,
	}

	//Keep reading data package from the client
	for {
		mes, err := msgTransformer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出..")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}

		// Process message
		// mes : Message
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
