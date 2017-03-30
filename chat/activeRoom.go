package chat

import "fmt"

type ActiveRoom struct {
	OnlineUsers map[string]*OnlineUser // 用户列表map
	Broadcast   chan Message           //消息channel
	CloseSign   chan bool              // 开关channel
}

type TextMessage struct {
	Content  string
	UserInfo *User
	Time     string
}

type UserStatus struct {
	Users []*User
}

type Message struct {
	MType       string
	TextMessage TextMessage
	UserStatus  UserStatus
}

//广播消息
func (_acRoom *ActiveRoom) broadcastUsers() {
	fmt.Println("broadcastUsers====")
	Users := runningActiveRoom.getOnlineUsers()
	m := Message{
		MType: STATUS_MTYPE,
		UserStatus: UserStatus{
			Users,
		},
	}

	runningActiveRoom.Broadcast <- m
}

//获取在线用户
func (_acRoom *ActiveRoom) getOnlineUsers() []*User {
	var users []*User
	for _, online := range _acRoom.OnlineUsers {
		users = append(users, online.UserInfo)
	}
	return users
}

//room初始化===
func (_acRoom *ActiveRoom) run() {
	for {
		select {
		case b := <-_acRoom.Broadcast:
			for _, online := range _acRoom.OnlineUsers {
				online.Send <- b
			}
		case c := <-_acRoom.CloseSign:
			if c == true {
				close(_acRoom.Broadcast)
				close(_acRoom.CloseSign)
				return
			}
		}
	}
}
