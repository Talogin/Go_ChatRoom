package process

import "fmt"

var (
	userMgr *UserManager
)

type UserManager struct {
	onlineUsers map[int]*UserRequestProcessor
}

func init() {
	userMgr = &UserManager{
		onlineUsers: make(map[int]*UserRequestProcessor, 1024),
	}
}

func (this *UserManager) AddOnlineUser(up *UserRequestProcessor) {
	this.onlineUsers[up.UserId] = up
}

func (this *UserManager) DeleteOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

func (this *UserManager) GetAllOnlineUsers() map[int]*UserRequestProcessor {
	return this.onlineUsers
}

func (this *UserManager) GetOnlineUserById(userId int) (up *UserRequestProcessor, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("User%d does not exist", userId)
		return
	}
	return
}
