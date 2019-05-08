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
	"github.com/ennoo/rivet/shunt"
	"github.com/ennoo/rivet/trans/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// LB 是否开启负载均衡
var LB = false

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
	request.call(context, method, remote, uri, nil)
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
	if LB {
		add, err := shunt.RunShunt(remote)
		if nil != err {
			request.result.Fail(err.Error())
			context.JSON(http.StatusOK, request.result)
		} else {
			// todo 请求协议判定
			remoteNew := strings.Join([]string{"http://", add.Host, ":", strconv.Itoa(add.Port)}, "")
			request.call(context, method, remoteNew, uri, callback)
		}
	} else {
		request.call(context, method, remote, uri, callback)
	}
}

// call 请求转发处理方案
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
func (request *Request) call(context *gin.Context, method string, remote string, uri string, callback func() *response.Result) {
	req := context.Request
	restTransHandler := RestTransHandler{
		RestHandler: RestHandler{
			RemoteServer: remote,
			URI:          uri,
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
	done(context, request, body, err, callback)
}

// done 请求转发处理结果
//
// 转发请求或降级回调
func done(context *gin.Context, request *Request, body []byte, err error, callback func() *response.Result) {
	if err != nil {
		request.result.Callback(callback, err)
	} else {
		log.Trans.Debug("body = " + string(body))
		err := json.Unmarshal(body, &request.result)
		if nil != err {
			request.result.Fail(err.Error())
		}
	}
	context.JSON(http.StatusOK, request.result)
}
