package service

import (
	"Go_ChatRoom/common/message"
	"Go_ChatRoom/server/model"
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *MyDao
)

type MyDao struct {
	Pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *MyDao) {

	userDao = &MyDao{
		Pool: pool,
	}
	return
}

func (this *MyDao) GetUserById(conn redis.Conn, userId int) (user *model.User, err error) {

	// Check if the user exist in DB
	userString, err := redis.String(conn.Do("HGet", "users", userId))

	if err != nil {
		if err == redis.ErrNil {
			err = model.ErrUserNotExist
		}
		return
	}

	// Unmarshal user and return
	user = &model.User{}
	err = json.Unmarshal([]byte(userString), user)
	if err != nil {
		fmt.Println("getUser unmarshal err = ", err)
		return
	}
	fmt.Println("Get the user from DB=%v", user)

	return
}

func (this *MyDao) AddUser(user *message.User) (err error) {

	conn := this.Pool.Get()
	defer conn.Close()
	// Check if userId already exist
	_, err = this.GetUserById(conn, user.UserId)
	if err != nil {
		err = model.ErrUserExist
		return
	}

	// Convert the instance to string
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	// Add it to Database
	_, err = conn.Do("HSet", "users", fmt.Sprintf("%d", user.UserId), string(data))
	if err != nil {
		return
	}
	return
}

func (this *MyDao) Login(userId int, pwd string) (user *model.User, err error) {

	// Get connection
	conn := this.Pool.Get()
	defer conn.Close()

	// Check if exist
	user, err = this.GetUserById(conn, userId)
	if err != nil {
		return
	}

	// Check the password
	if user.UserPwd != pwd {
		err = model.ErrInvalidPassword
		return
	}

	return
}
