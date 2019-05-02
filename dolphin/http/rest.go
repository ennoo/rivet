/*
 * Copyright (c) 2019. ENNOO - All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
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
	remoteServer string
	uri          string
	param        interface{}
	values       url.Values
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
