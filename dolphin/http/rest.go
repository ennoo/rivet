package http

import (
	"github.com/ennoo/rivet/common/log"
	"github.com/ennoo/rivet/common/string"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

// RestHandler 处理请求发送和接收
type RestHandler struct {
	// 远程服务器地址,如 http://localhost:3030
	RemoteServer string
	uri          string
	param        interface{}
	values       url.Values
}

type JSONRestHandler struct {
	// 远程服务器地址,如 http://localhost:3030
	RemoteServer string
}

// 从Header或者Cookie中获取到用户的access_token
func GetAccessTokenFromReq(c *gin.Context) (token string) {
	var err error

	token = c.GetHeader("Access_token")
	if str.IsEmpty(token) {
		token = c.GetHeader("access_token")
		// 如果依然为空，则从cookie中尝试获取
		if str.IsEmpty(token) {
			token, err = c.Cookie("access_token")
			if err != nil {
				log.Error(err)
				return ""
			}
		}
	}
	return token
}

// 忽略大小写，找到指定的cookie
func GetCookieByName(cookies []*http.Cookie, cookieName string) *http.Cookie {
	for _, cookie := range cookies {
		if strings.ToLower(cookie.Name) == strings.ToLower(cookieName) {
			return cookie
		}
	}
	return nil
}
