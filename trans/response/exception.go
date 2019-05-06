package response

import (
	"net/http"
	"strings"
)

var (
	ExpNotExist = Exp("not exist", http.StatusOK)
)

type Exception struct {
	Msg  string // 异常通用信息
	code int    // http 请求返回 code
}

func Exp(brief string, httpCode int) Exception {
	return Exception{
		Msg:  brief,
		code: httpCode}
}

func (exception *Exception) Fit(prefix string) *Exception {
	exception.Msg = strings.Join([]string{prefix, exception.Msg}, " ")
	return exception
}
