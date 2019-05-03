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
	"bytes"
	"encoding/json"
	"github.com/ennoo/rivet/common/util/log"
	"github.com/ennoo/rivet/dolphin/http/response"
	"net/http"
	"net/url"
	"path/filepath"
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

	ObtainParam() interface{}

	ObtainValue() url.Values

	ObtainRemoteServer() string

	ObtainHeader() http.Header

	ObtainCookies() []*http.Cookie
}

func requestJson(method string, handler Handler) ([]byte, error) {
	// 将参数转化为json比特流
	jsonByte, _ := json.Marshal(handler.ObtainParam())
	log.Debug("body:", string(jsonByte))
	req, err := http.NewRequest(method, getFullUri(handler), bytes.NewReader(jsonByte))
	if nil != err {
		return nil, err
	}
	req.Cookies() = handler.ObtainCookies()
	req.Header = handler.ObtainHeader()
	return response.Response(req)
}

func requestText(method string, handler Handler) ([]byte, error) {
	req, err := http.NewRequest(method, getFullUri(handler), bytes.NewBufferString(handler.ObtainValue().Encode()))
	if nil != err {
		return nil, err
	}
	req.Cookies() = handler.ObtainCookies()
	req.Header = handler.ObtainHeader()
	return response.Response(req)
}

// Get 发送get请求
func get(handler Handler) (body []byte, err error) {
	// 将参数转化为json比特流
	jsonByte, _ := json.Marshal(handler.ObtainParam())
	log.Debug("body:", string(jsonByte))
	req, err := http.NewRequest(http.MethodGet, getFullUri(handler), nil)
	if nil != err {
		return nil, err
	}
	req.Cookies() = handler.ObtainCookies()
	req.Header = handler.ObtainHeader()
	return response.Response(req)
}

func getFullUri(handler Handler) string {
	return filepath.ToSlash(filepath.Join(handler.ObtainRemoteServer(), "/", handler.ObtainUri()))
}
