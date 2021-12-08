package model

import (
	"Go_ChatRoom/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
