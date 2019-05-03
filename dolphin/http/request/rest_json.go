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
	"io"
	"net/http"
)

type RestJsonHandler struct {
	RestHandler
	Param interface{}
}

func (handler *RestJsonHandler) ObtainUri() string {
	return handler.Uri
}

func (handler *RestJsonHandler) ObtainRemoteServer() string {
	return handler.RemoteServer
}

func (handler *RestJsonHandler) ObtainBody() io.Reader {
	jsonByte, _ := json.Marshal(handler.Param)
	return bytes.NewReader(jsonByte)
}

func (handler *RestJsonHandler) ObtainHeader() http.Header {
	handler.Header.Add("Content-Type", "application/json")
	return handler.Header
}

func (handler *RestJsonHandler) ObtainCookies() []http.Cookie {
	return handler.Cookies
}

func (handler *RestJsonHandler) Post() (body []byte, err error) {
	return request(http.MethodPost, handler)
}

func (handler *RestJsonHandler) Put() (body []byte, err error) {
	return request(http.MethodPut, handler)
}

func (handler *RestJsonHandler) Delete() (body []byte, err error) {
	return request(http.MethodDelete, handler)
}

func (handler *RestJsonHandler) Patch() (body []byte, err error) {
	return request(http.MethodPatch, handler)
}

func (handler *RestJsonHandler) Options() (body []byte, err error) {
	return request(http.MethodOptions, handler)
}

func (handler *RestJsonHandler) Head() (body []byte, err error) {
	return request(http.MethodHead, handler)
}

func (handler *RestJsonHandler) Connect() (body []byte, err error) {
	return request(http.MethodConnect, handler)
}

func (handler *RestJsonHandler) Trace() (body []byte, err error) {
	return request(http.MethodTrace, handler)
}

// Get 发送get请求
func (handler *RestJsonHandler) Get() (body []byte, err error) {
	return get(handler)
}
