package response

import "time"

type Response struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
	Timestamp int64  `json:"timestamp"`
}

func Success() *Response {
	return &Response{
		Code:      200,
		Message:   "请求成功",
		Data:      nil,
		Timestamp: time.Now().Unix(),
	}
}

func SuccessWithData(data any) *Response {
	return &Response{
		Code:      200,
		Message:   "请求成功",
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
}

func SuccessWithMessage(message string) *Response {
	return &Response{
		Code:      200,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	}
}

func Fail() *Response {
	return &Response{
		Code:      500,
		Message:   "请求失败",
		Data:      nil,
		Timestamp: time.Now().Unix(),
	}
}

func FailWithMessage(message string) *Response {
	return &Response{
		Code:      500,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	}
}

func FailWithCodeAndMessage(code int, message string) *Response {
	return &Response{
		Code:      code,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	}
}
