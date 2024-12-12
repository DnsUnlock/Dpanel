package response

import "github.com/DnsUnlock/Dpanel/backend/model/response/statusPrefix"

type Response struct {
	Code string      `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Err  string      `json:"err,omitempty"`
}

func Success(Msg string, data interface{}) Response {
	return Response{
		Code: statusPrefix.OK.String(),
		Data: data,
		Msg:  Msg,
	}
}

func NotFound(Msg string, data interface{}) *Response {
	return &Response{
		Code: statusPrefix.NotFound.String(),
		Data: data,
		Msg:  Msg,
	}
}
func Error(Err string, data interface{}) *Response {
	return &Response{
		Code: statusPrefix.ERROR.String(),
		Data: data,
		Err:  Err,
	}
}
