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
	"net/http"
	"net/url"
)

type RestTextHandler struct {
	RestHandler
	Header  http.Header
	Cookies []http.Cookie
}

func (handler *RestTextHandler) ObtainUri() string {
	return handler.RestHandler.Uri
}

func (handler *RestTextHandler) ObtainParam() interface{} {
	return nil
}

func (handler *RestTextHandler) ObtainValue() url.Values {
	return handler.Values
}

func (handler *RestTextHandler) ObtainRemoteServer() string {
	return handler.RestHandler.RemoteServer
}

func (handler *RestTextHandler) ObtainHeader() http.Header {
	handler.Header.Add("Content-Type", "text/html")
	return handler.Header
}

func (handler *RestTextHandler) ObtainCookies() []http.Cookie {
	return handler.Cookies
}

func (handler *RestTextHandler) Post() (body []byte, err error) {
	return requestText(http.MethodPost, handler)
}

func (handler *RestTextHandler) Put() (body []byte, err error) {
	return requestText(http.MethodPut, handler)
}

func (handler *RestTextHandler) Delete() (body []byte, err error) {
	return requestText(http.MethodDelete, handler)
}

func (handler *RestTextHandler) Patch() (body []byte, err error) {
	return requestText(http.MethodPatch, handler)
}

func (handler *RestTextHandler) Options() (body []byte, err error) {
	return requestText(http.MethodOptions, handler)
}

func (handler *RestTextHandler) Head() (body []byte, err error) {
	return requestText(http.MethodHead, handler)
}

func (handler *RestTextHandler) Connect() (body []byte, err error) {
	return requestText(http.MethodConnect, handler)
}

func (handler *RestTextHandler) Trace() (body []byte, err error) {
	return requestText(http.MethodTrace, handler)
}

// Get 发送get请求
func (handler *RestTextHandler) Get() (body []byte, err error) {
	return get(handler)
}
