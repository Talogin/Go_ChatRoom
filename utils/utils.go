package utils

import (
	"Go_ChatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type MessageTransformer struct {
	Conn   net.Conn
	Buffer [8096]byte
}

func (this *MessageTransformer) WritePkg(mes []byte) (err error) {

	//Send length of data to the server
	//1. Calculate the length of data
	packageLength := uint32(len(mes))

	//2. Convert unit32 to byte[]
	binary.BigEndian.PutUint32(this.Buffer[0:4], packageLength)

	//3. Send
	n, err := this.Conn.Write(this.Buffer[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//Send message data
	n, err = this.Conn.Write(mes)
	if n != int(packageLength) || err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	return
}
func (this *MessageTransformer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("Reading the data from client...")

	//Read package length
	_, err = this.Conn.Read(this.Buffer[:4])
	if err != nil {
		return
	}

	packageLength := binary.BigEndian.Uint32(this.Buffer[:4])

	//Read message data
	n, err := this.Conn.Read(this.Buffer[:packageLength])
	if n != int(packageLength) || err != nil {
		return
	}

	err = json.Unmarshal(this.Buffer[:packageLength], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return

}
