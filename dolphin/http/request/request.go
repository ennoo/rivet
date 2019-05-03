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
	"github.com/ennoo/rivet/dolphin/http/response"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type Request interface {
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
	return response.Response(req)
}

// Get 发送get请求
func get(handler Handler) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, getFullUri(handler), nil)
	if nil != err {
		return nil, err
	}
	addCookies(req, handler.ObtainCookies())
	req.Header = handler.ObtainHeader()
	return response.Response(req)
}

func getFullUri(handler Handler) string {
	return filepath.ToSlash(strings.Join([]string{handler.ObtainRemoteServer(), filepath.Join("/", handler.ObtainUri())}, ""))
}
