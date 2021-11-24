package service

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type DBContext struct {
	Conn redis.Conn
}

func(this *DBContext) getUserById(userId int, pwd string) (err error) {
	if(userId == nil || pwd == nil) {
		fmt.Println("UserId or password you entered cannot be empty...")
		return
	}


}