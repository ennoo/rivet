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
package request

import (
	"github.com/ennoo/rivet/common/util/log"
	"github.com/ennoo/rivet/common/util/string"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

// RestHandler 处理请求发送和接收
type RestHandler struct {
	// 远程服务器地址,如 http://localhost:3030
	RemoteServer string
	Uri          string
	Body         io.ReadCloser
	Header       http.Header
	Cookies      []http.Cookie
}

type Rest interface {
	Post() (body []byte, err error)

	Put() (body []byte, err error)

	Delete() (body []byte, err error)

	Patch() (body []byte, err error)

	Options() (body []byte, err error)

	Head() (body []byte, err error)

	Connect() (body []byte, err error)

	Trace() (body []byte, err error)

	Get() (body []byte, err error)
}

type Handler interface {
	ObtainUri() string

	ObtainRemoteServer() string

	ObtainBody() io.Reader

	ObtainHeader() http.Header

	ObtainCookies() []http.Cookie
}

func addCookies(request *http.Request, cookies []http.Cookie) {
	for _, cookie := range cookies {
		request.AddCookie(&cookie)
	}
}

func request(method string, handler Handler) ([]byte, error) {
	req, err := http.NewRequest(method, getFullUri(handler), handler.ObtainBody())
	if nil != err {
		return nil, err
	}
	addCookies(req, handler.ObtainCookies())
	req.Header = handler.ObtainHeader()
	return exec(req)
}

// Get 发送get请求
func get(handler Handler) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, getFullUri(handler), nil)
	if nil != err {
		return nil, err
	}
	addCookies(req, handler.ObtainCookies())
	req.Header = handler.ObtainHeader()
	return exec(req)
}

func exec(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func getFullUri(handler Handler) string {
	return filepath.ToSlash(strings.Join([]string{handler.ObtainRemoteServer(), filepath.Join("/", handler.ObtainUri())}, ""))
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
