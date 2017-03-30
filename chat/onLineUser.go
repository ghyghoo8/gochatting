package chat

import (
	"github.com/gorilla/websocket"
	"time"
	"fmt"
)

type OnlineUser struct {
	InRoom     *ActiveRoom
	Connection *websocket.Conn
	UserInfo   *User
	Send       chan Message
}

//用户下线
func (_user *OnlineUser) killUserResource(){
	//关闭连接===
	fmt.Println("killUserResource===")
	_user.Connection.Close()
	_room:=_user.InRoom
	delete(_room.OnlineUsers, _user.UserInfo.Email)
	close(_user.Send)
	runningActiveRoom.broadcastUsers()
}

//获取客户端消息
func (_fromClient *OnlineUser) PullFromClient() {
	fmt.Println("PullFromClient===")
	conn:=_fromClient.Connection
	//死循环=== 除非遇到断言，return、break。。。
	for {

		//content是[]btye 类型
		msgType, content, err := conn.ReadMessage()
		fmt.Println("===recive:",string(content),msgType)
		if err != nil {
			fmt.Println("out for====")
			return
		}

		m := Message{
			MType: TEXT_MTYPE,
			TextMessage: TextMessage{
				UserInfo: _fromClient.UserInfo,
				Time:     humanCreatedAt(),
				Content:  string(content),
			},
		}
		fmt.Println(m)
		_fromClient.InRoom.Broadcast <- m
	}
}

//push消息到客户端
func (_toClient *OnlineUser) PushToClient() {
	fmt.Println("PushToClient===")
	for b := range _toClient.Send {
		err := websocket.WriteJSON(_toClient.Connection, b)
		if err != nil {
			break
		}
	}
}

//格式化登陆时间
func humanCreatedAt() string {
	return time.Now().Format(TIME_FORMAT)
}