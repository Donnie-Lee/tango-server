package websocket

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"imserver/models/chatModel"
)

func Send(request any) {

	var chatMessage *chatModel.ChatMessage
	err := mapstructure.Decode(request, &chatMessage)
	if err != nil {
		fmt.Println("请求参数解析失败")
		return
	}

	WebSocketManage.Send(chatMessage)
}
