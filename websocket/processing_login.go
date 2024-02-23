package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"imserver/common/response"
)

type LoginRequest struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(request any, ws *websocket.Conn) *response.Response {

	var loginRequest *LoginRequest
	err := mapstructure.Decode(request, &loginRequest)
	if err != nil {
		return response.FailWithMessage("请求参数解析失败")
	}
	if loginRequest.Password == "123456" || loginRequest.Password == "admin" {
		WebSocketManage.Register(&UserCon{
			UserId: loginRequest.UserId,
			Con:    ws,
		})
		return response.SuccessWithMessage("登录成功")
	}

	return response.FailWithMessage("登陆失败")

}
