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

// RestHandler 处理请求发送内容
type RestHandler struct {
	RemoteServer string
	URI          string
	Body         io.ReadCloser
	Header       http.Header
	Cookies      []http.Cookie
}

// Rest http 请求方法接口
type Rest interface {
	// Post 发起 Post 请求，body 为请求后的返回内容，err 指出请求出错原因
	Post() (body []byte, err error)

	// Put 发起 Put 请求，body 为请求后的返回内容，err 指出请求出错原因
	Put() (body []byte, err error)

	// Delete 发起 Delete 请求，body 为请求后的返回内容，err 指出请求出错原因
	Delete() (body []byte, err error)

	// Patch 发起 Patch 请求，body 为请求后的返回内容，err 指出请求出错原因
	Patch() (body []byte, err error)

	// Options 发起 Options 请求，body 为请求后的返回内容，err 指出请求出错原因
	Options() (body []byte, err error)

	// Head 发起 Head 请求，body 为请求后的返回内容，err 指出请求出错原因
	Head() (body []byte, err error)

	// Connect 发起 Connect 请求，body 为请求后的返回内容，err 指出请求出错原因
	Connect() (body []byte, err error)

	// Trace 发起 Trace 请求，body 为请求后的返回内容，err 指出请求出错原因
	Trace() (body []byte, err error)

	// Get 发起 Get 请求，body 为请求后的返回内容，err 指出请求出错原因
	Get() (body []byte, err error)
}

// Handler http 处理请求发送内容接口
type Handler interface {
	// ObtainRemoteServer 获取本次 http 请求服务根路径 如：localhost:8080
	ObtainRemoteServer() string

	// ObtainURI 获取本次 http 请求服务方法路径 如：/user/login
	ObtainURI() string

	// ObtainBody 获取本次 http 请求 body io
	ObtainBody() io.Reader

	// ObtainHeader 获取本次 http 请求 header
	ObtainHeader() http.Header

	// ObtainCookies 获取本次 http 请求 cookies
	ObtainCookies() []http.Cookie
}

func addCookies(request *http.Request, cookies []http.Cookie) {
	for _, cookie := range cookies {
		request.AddCookie(&cookie)
	}
}

func request(method string, handler Handler) ([]byte, error) {
	req, err := http.NewRequest(method, getFullURI(handler), handler.ObtainBody())
	if nil != err {
		return nil, err
	}
	addCookies(req, handler.ObtainCookies())
	req.Header = handler.ObtainHeader()
	return exec(req)
}

// Get 发送get请求
func get(handler Handler) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, getFullURI(handler), nil)
	if nil != err {
		return nil, err
	}
	addCookies(req, handler.ObtainCookies())
	req.Header = handler.ObtainHeader()
	return exec(req)
}

func exec(req *http.Request) ([]byte, error) {
	client := http.Client{
		Transport: GetTPInstance().Transport,
	}
	resp, err := client.Do(req)

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

func getFullURI(handler Handler) string {
	return filepath.ToSlash(strings.Join([]string{handler.ObtainRemoteServer(), filepath.Join("/", handler.ObtainURI())}, ""))
}

// GetAccessTokenFromReq 从Header或者Cookie中获取到用户的access_token
func GetAccessTokenFromReq(c *gin.Context) (token string) {
	var err error

	token = c.GetHeader("Access_token")
	if str.IsEmpty(token) {
		token = c.GetHeader("access_token")
		// 如果依然为空，则从cookie中尝试获取
		if str.IsEmpty(token) {
			token, err = c.Cookie("access_token")
			if err != nil {
				log.Trans.Error(err.Error())
				return ""
			}
		}
	}
	return token
}

// GetCookieByName 忽略大小写，找到指定的cookie
func GetCookieByName(cookies []*http.Cookie, cookieName string) *http.Cookie {
	for _, cookie := range cookies {
		if strings.EqualFold(cookie.Name, cookieName) {
			return cookie
		}
	}
	return nil
}
