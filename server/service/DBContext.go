package service

import (
	"context"
	"fmt"
)

type DBContext struct{}

func (this *DBContext) getUserById(userId int, pwd string) (err error) {

	ctx := context.Background()
	// Get redis connection
	conn := this.Pool.Get(ctx)
	defer conn.Close()

	if userId == 0 || pwd == "" {
		fmt.Println("UserId or password you entered cannot be empty...")
		return
	}

}
