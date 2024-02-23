package chatModel

type ChatMessage struct {
	// 0 单聊 1 群聊
	ChatRoomId int    `json:"chatRoomId"`
	SenderId   string `json:"senderId"`
	Message    string `json:"message"`
	ReceiverId string `json:"receiverId"`
	// 0 文本消息 1 语音 2 视频 3 图片
	MessageType int    `json:"messageType"`
	Canceled    int    `json:"canceled"`
	SendTime    string `json:"sendTime"`
	QuoteId     int    `json:"quoteId"`
	Status      int    `json:"status"`
}
