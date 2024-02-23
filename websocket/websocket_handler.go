package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"imserver/common/request"
)

type WebsocketHandler struct {
}

func NewWebsocketHandler() *WebsocketHandler {
	return &WebsocketHandler{}
}

func (wh *WebsocketHandler) Handler(c *gin.Context) {
	// 获取WebSocket连接
	ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		panic(err)
	}

	// 处理WebSocket消息
	go wh.handleRead(ws, c)
}

func (wh *WebsocketHandler) handleRead(ws *websocket.Conn, c *gin.Context) {
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			WebSocketManage.Cancel(ws)
			return
		}

		var request *request.Request
		err = json.Unmarshal(message, &request)
		if err != nil {
			fmt.Println("请求信息解析失败", err)
			continue
		}

		var response any
		switch request.Header.ProcessMethod {
		case "loginModel":
			Login(request.Body, ws)
			break
		case "send":
			Send(request.Body)
			break
		case "heartBeat":
			break
		}

		if response != nil {
			ws.WriteJSON(response)
		}
	}
}
