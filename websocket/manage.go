package websocket

import (
	"github.com/gorilla/websocket"
	"imserver/middleware/log"
	"imserver/models/chatModel"
)

type WebSocketMap struct {
	UserMap    map[string]*UserCon
	ReflectMap map[*websocket.Conn]string
}

type UserCon struct {
	UserId string
	Con    *websocket.Conn
}

var WebSocketManage *WebSocketMap

func init() {
	WebSocketManage = &WebSocketMap{
		UserMap:    make(map[string]*UserCon),
		ReflectMap: make(map[*websocket.Conn]string),
	}
}

func (wm *WebSocketMap) Register(userCon *UserCon) {
	log.Logger.Infof("注册通道 %s", userCon.UserId)
	wm.UserMap[userCon.UserId] = userCon
	wm.ReflectMap[userCon.Con] = userCon.UserId
}

func (wm *WebSocketMap) Cancel(con *websocket.Conn) {
	userId := wm.ReflectMap[con]
	log.Logger.Info("断开通道 %s", userId)
	delete(wm.UserMap, userId)
	delete(wm.ReflectMap, con)
	con.Close()
}

func (wm *WebSocketMap) Send(chatMessage *chatModel.ChatMessage) {
	if wm.UserMap[chatMessage.ReceiverId] != nil {
		wm.UserMap[chatMessage.ReceiverId].Con.WriteJSON(chatMessage)
	}
}

func (wm *WebSocketMap) SendByGroup(chatMessage *chatModel.ChatMessage) {
	//for index := range chatMessage.GroupTo {
	//	chatMessage.To = chatMessage.GroupTo[index]
	//	wm.Send(chatMessage)
	//}
}
