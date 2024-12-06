package common

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
}

func Response(w http.ResponseWriter, code int, message string, data interface{}) {
	httpx.OkJson(w, Body{
		Code:    code,
		Message: message,
		Result:  data,
	})
}

// Success 成功的请求
func Success(w http.ResponseWriter, data interface{}) {
	Response(w, 0, "请求成功", data)
}

// Fail 失败的请求
func Fail(w http.ResponseWriter, message string) {
	Response(w, 1, message, nil)
}
