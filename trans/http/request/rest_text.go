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
	"io"
	"net/http"
	"net/url"
)

type RestTextHandler struct {
	RestHandler
	Values url.Values
}

func (handler *RestTextHandler) ObtainUri() string {
	return handler.RestHandler.Uri
}

func (handler *RestTextHandler) ObtainRemoteServer() string {
	return handler.RemoteServer
}

func (handler *RestTextHandler) ObtainBody() io.Reader {
	return bytes.NewBufferString(handler.Values.Encode())
}

func (handler *RestTextHandler) ObtainHeader() http.Header {
	handler.Header.Add("Content-Type", "text/html")
	return handler.Header
}

func (handler *RestTextHandler) ObtainCookies() []http.Cookie {
	return handler.Cookies
}

func (handler *RestTextHandler) Post() (body []byte, err error) {
	return request(http.MethodPost, handler)
}

func (handler *RestTextHandler) Put() (body []byte, err error) {
	return request(http.MethodPut, handler)
}

func (handler *RestTextHandler) Delete() (body []byte, err error) {
	return request(http.MethodDelete, handler)
}

func (handler *RestTextHandler) Patch() (body []byte, err error) {
	return request(http.MethodPatch, handler)
}

func (handler *RestTextHandler) Options() (body []byte, err error) {
	return request(http.MethodOptions, handler)
}

func (handler *RestTextHandler) Head() (body []byte, err error) {
	return request(http.MethodHead, handler)
}

func (handler *RestTextHandler) Connect() (body []byte, err error) {
	return request(http.MethodConnect, handler)
}

func (handler *RestTextHandler) Trace() (body []byte, err error) {
	return request(http.MethodTrace, handler)
}

// Get 发送get请求
func (handler *RestTextHandler) Get() (body []byte, err error) {
	return get(handler)
}
