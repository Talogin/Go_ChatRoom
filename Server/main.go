package main

import (
	"fmt"
)

func main()  {
	
	//Listen to port 8080
	In, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Get errors when trying to listen to port 8080")
	}
	for {
		conn, err := In.Accept()
		if err != nil {
			fmt.Println("")
		}
		go 
	}
}