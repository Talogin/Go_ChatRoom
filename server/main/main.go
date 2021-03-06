package main

import (
	"Go_ChatRoom/server/service"
	"fmt"
	"net"
	"time"

	"github.com/gomodule/redigo/redis"
)

func connectionProcess(conn net.Conn) {

	userConnectionHandler := &UserConnectionHandler{
		Conn: conn,
	}
	err := userConnectionHandler.process()
	if err != nil {
		fmt.Println("Communication between server and client=err", err)
		return
	}
	defer conn.Close()
}

// Initializatin of redis pool
func initPool() {
	Pool = &redis.Pool{
		MaxIdle:     2, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   3, //最大数
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH", "12345"); err != nil {
				c.Close()
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func initUserDao() {
	service.MyUserDao = service.NewUserDao(Pool)
}

func init() {

	initPool()
	initUserDao()
}

var Pool *redis.Pool

func main() {

	//Listen to port 8889
	fmt.Println("The server is listening to port 8889")
	In, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("Get errors when trying to listen to port 8889")
	}

	//Initialize redis pool
	initPool()

	for {
		conn, err := In.Accept()
		if err != nil {
			fmt.Println("Building net connection failed...")
		}
		go connectionProcess(conn)
	}
}
