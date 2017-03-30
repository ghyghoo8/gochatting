package chat

import (
	"strings"
	"github.com/gorilla/websocket"
	"../libs"
	"net/http"
	"log"
	"fmt"
)

var (
	//全局room
	runningActiveRoom *ActiveRoom = &ActiveRoom{}
	//设置websocket参数===
	 upgrader = websocket.Upgrader{}
)

//初始化聊天室===
func InitChatRoom() {
	fmt.Println("start a chatRoom=====")
	//初始化实例===
	runningActiveRoom = &ActiveRoom{
		OnlineUsers: make(map[string]*OnlineUser),
		Broadcast:   make(chan Message),
		CloseSign:   make(chan bool),
	}
	//起一个
	go runningActiveRoom.run()
}

//建立连接。。。ws://
func BuildConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	email := r.URL.Query().Get("email")
	if email == "" {
		email = "ghyghoo8@gmail.com"
	}
	onlineUser := &OnlineUser{
		InRoom:     runningActiveRoom,
		Connection: ws,
		Send:       make(chan Message, 256),
		UserInfo: &User{
			Email:    email,
			Name:     strings.Split(email, "@")[0],
			Gravatar: libs.UrlSize(email, 20),
		},
	}
	runningActiveRoom.OnlineUsers[email] = onlineUser

	runningActiveRoom.broadcastUsers()


	go onlineUser.PushToClient()
	onlineUser.PullFromClient()

	defer onlineUser.killUserResource()

}
