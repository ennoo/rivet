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
	"encoding/json"
	"github.com/ennoo/rivet/common/util/log"
	"github.com/ennoo/rivet/trans/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Request 提供实例化调用请求方法，并内置返回策略
type Request struct {
	result response.Result
}

// Call 请求转发处理方案
//
// context：原请求上下文
//
// method：即将转发的请求方法
//
// remote：请求转发主体域名
//
// uri：请求转发主体方法路径
func (request *Request) Call(context *gin.Context, method string, remote string, uri string) {
	request.Callback(context, method, remote, uri, nil)
}

// Callback 请求转发处理方案
//
// context：原请求上下文
//
// method：即将转发的请求方法
//
// remote：请求转发主体域名
//
// uri：请求转发主体方法路径
//
// callback：请求转发失败后回调降级策略
//
// callback *response.Result 请求转发降级后返回请求方结果对象
func (request *Request) Callback(context *gin.Context, method string, remote string, uri string, callback func() *response.Result) {
	req := context.Request
	restTransHandler := RestTransHandler{
		RestHandler: RestHandler{
			RemoteServer: remote,
			Uri:          uri,
			Body:         req.Body,
			Header:       nil,
			Cookies:      nil}}
	var body []byte
	var err error

	switch method {
	case http.MethodGet:
		body, err = restTransHandler.Get()
	case http.MethodHead:
		body, err = restTransHandler.Head()
	case http.MethodPost:
		body, err = restTransHandler.Post()
	case http.MethodPut:
		body, err = restTransHandler.Put()
	case http.MethodPatch:
		body, err = restTransHandler.Patch()
	case http.MethodDelete:
		body, err = restTransHandler.Delete()
	case http.MethodConnect:
		body, err = restTransHandler.Connect()
	case http.MethodOptions:
		body, err = restTransHandler.Options()
	case http.MethodTrace:
		body, err = restTransHandler.Trace()
	}
	if err != nil {
		request.result.Callback(callback, err)
	} else {
		log.Debug("body = ", string(body))
		err := json.Unmarshal(body, &request.result)
		if nil != err {
			request.result.Fail(err.Error())
		}
	}
	context.JSON(http.StatusOK, request.result)
}
