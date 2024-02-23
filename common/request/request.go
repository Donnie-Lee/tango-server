package request

type Request struct {
	Header *RequestHeader `json:"header"`
	Body   any            `json:"body"`
}

type RequestHeader struct {
	// 消息类型 0 字符 1 字节
	MessageType int `json:"messageType"`
	// 用户标识
	Authorization string `json:"authorization"`
	// 业务处理方法
	ProcessMethod string `json:"processMethod"`
}

func NewRequest(header *RequestHeader, body any) *Request {
	return &Request{
		Header: header,
		Body:   body,
	}
}

func (r *Request) GetHeader() *RequestHeader {
	return r.Header
}

func (r *Request) GetBody() any {
	return r.Body
}
