package main

import (
	"fmt"
	"net"
	"Go_ChatRoom/utils"
)

type UserConnectionHandler struct {
	Conn net.Conn
}

func (this *UserConnectionHandler)process() (err error){

	msgTransformer := &utils.MessageTransformer {
		Conn : this.Conn,
	}
	//Keep reading data package from the client
	for {
		mes, err := msgTransformer.ReadPkg()
		if err != nil {
			fmt.Println("readPkg err=", err)
			return err
		}
		fmt.Println(mes)
	}
}