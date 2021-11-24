package main

import (
	"fmt"
	"net"
	"github.com/go-redis/redis/v8"
)

func process(conn net.Conn) {

	defer conn.Close()
	
	userConnectionHandler := &UserConnectionHandler{
		Conn : conn,
	}
	err := userConnectionHandler.process()
	if err != nil {
		fmt.Println("Communication between server and client=err", err)
		return 
	}
}

// Initializatin of redis pool
func init(){
    pool = &redis.Pool{     
        MaxIdle:16,    
        MaxActive:0,    
        IdleTimeout:300,        
        Dial: func() (redis.Conn ,error){  
            return redis.Dial("tcp","localhost:6379")
        },
    }
}

var pool *redis.Pool

func main()  {

	
	//Listen to port 8889
	fmt.Println("The server is listening to port 8889")
	In, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("Get errors when trying to listen to port 8889")
	}

	//Initialize redis pool
	init()

	for {
		conn, err := In.Accept()
		if err != nil {
			fmt.Println("Building net connection failed...")
		}
		go process(conn) 
	}
}