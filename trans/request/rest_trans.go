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
	"io"
	"net/http"
)

type RestTransHandler struct {
	RestHandler
}

func (handler *RestTransHandler) ObtainUri() string {
	return handler.RestHandler.Uri
}

func (handler *RestTransHandler) ObtainRemoteServer() string {
	return handler.RemoteServer
}

func (handler *RestTransHandler) ObtainBody() io.Reader {
	return handler.Body
}

func (handler *RestTransHandler) ObtainHeader() http.Header {
	return handler.Header
}

func (handler *RestTransHandler) ObtainCookies() []http.Cookie {
	return handler.Cookies
}

func (handler *RestTransHandler) Post() (body []byte, err error) {
	return request(http.MethodPost, handler)
}

func (handler *RestTransHandler) Put() (body []byte, err error) {
	return request(http.MethodPut, handler)
}

func (handler *RestTransHandler) Delete() (body []byte, err error) {
	return request(http.MethodDelete, handler)
}

func (handler *RestTransHandler) Patch() (body []byte, err error) {
	return request(http.MethodPatch, handler)
}

func (handler *RestTransHandler) Options() (body []byte, err error) {
	return request(http.MethodOptions, handler)
}

func (handler *RestTransHandler) Head() (body []byte, err error) {
	return request(http.MethodHead, handler)
}

func (handler *RestTransHandler) Connect() (body []byte, err error) {
	return request(http.MethodConnect, handler)
}

func (handler *RestTransHandler) Trace() (body []byte, err error) {
	return request(http.MethodTrace, handler)
}

// Get 发送get请求
func (handler *RestTransHandler) Get() (body []byte, err error) {
	return get(handler)
}
