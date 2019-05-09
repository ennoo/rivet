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
	"net/url"
	"strconv"
	"strings"
	"sync"
)

var (
	// LB 是否开启负载均衡
	LB  = false
	req = sync.Pool{
		New: func() interface{} {
			return &Request{}
		},
	}
)

// Request 提供实例化调用请求方法，并内置返回策略
type Request struct {
	result response.Result
}

// SyncPoolGetRequest 提供实例化调用请求方法，并内置返回策略
func SyncPoolGetRequest() *Request {
	return req.Get().(*Request)
}

// RestJson JSON 请求
//
// method：请求方法
//
// remote：请求主体域名
//
// uri：请求主体方法路径
//
// param 请求对象
func (request *Request) RestJson(method string, remote string, uri string, param interface{}) ([]byte, error) {
	restJSONHandler := RestJSONHandler{
		RestHandler: RestHandler{
			RemoteServer: remote,
			URI:          uri,
		},
		Param: param,
	}
	var body []byte
	var err error

	switch method {
	case http.MethodGet:
		body, err = restJSONHandler.Get(DirectJSONRequest)
	case http.MethodHead:
		body, err = restJSONHandler.Head(DirectJSONRequest)
	case http.MethodPost:
		body, err = restJSONHandler.Post(DirectJSONRequest)
	case http.MethodPut:
		body, err = restJSONHandler.Put(DirectJSONRequest)
	case http.MethodPatch:
		body, err = restJSONHandler.Patch(DirectJSONRequest)
	case http.MethodDelete:
		body, err = restJSONHandler.Delete(DirectJSONRequest)
	case http.MethodConnect:
		body, err = restJSONHandler.Connect(DirectJSONRequest)
	case http.MethodOptions:
		body, err = restJSONHandler.Options(DirectJSONRequest)
	case http.MethodTrace:
		body, err = restJSONHandler.Trace(DirectJSONRequest)
	}
	return body, err
}

// RestText TEXT 请求
//
// method：请求方法
//
// remote：请求主体域名
//
// uri：请求主体方法路径
//
// values 请求参数
func (request *Request) RestText(method string, remote string, uri string, values url.Values) ([]byte, error) {
	restTextHandler := RestTextHandler{
		Values: values,
	}
	var body []byte
	var err error

	switch method {
	case http.MethodGet:
		body, err = restTextHandler.Get(DirectTextRequest)
	case http.MethodHead:
		body, err = restTextHandler.Head(DirectTextRequest)
	case http.MethodPost:
		body, err = restTextHandler.Post(DirectTextRequest)
	case http.MethodPut:
		body, err = restTextHandler.Put(DirectTextRequest)
	case http.MethodPatch:
		body, err = restTextHandler.Patch(DirectTextRequest)
	case http.MethodDelete:
		body, err = restTextHandler.Delete(DirectTextRequest)
	case http.MethodConnect:
		body, err = restTextHandler.Connect(DirectTextRequest)
	case http.MethodOptions:
		body, err = restTextHandler.Options(DirectTextRequest)
	case http.MethodTrace:
		body, err = restTextHandler.Trace(DirectTextRequest)
	}
	return body, err
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
	cookies := req.Cookies()
	restTransHandler := RestTransHandler{
		RestHandler: RestHandler{
			RemoteServer: remote,
			URI:          uri,
			Body:         req.Body,
			Header:       req.Header,
			Cookies:      cookies}}
	var body []byte
	var err error

	switch method {
	case http.MethodGet:
		body, err = restTransHandler.Get(TransCallbackRequest)
	case http.MethodHead:
		body, err = restTransHandler.Head(TransCallbackRequest)
	case http.MethodPost:
		body, err = restTransHandler.Post(TransCallbackRequest)
	case http.MethodPut:
		body, err = restTransHandler.Put(TransCallbackRequest)
	case http.MethodPatch:
		body, err = restTransHandler.Patch(TransCallbackRequest)
	case http.MethodDelete:
		body, err = restTransHandler.Delete(TransCallbackRequest)
	case http.MethodConnect:
		body, err = restTransHandler.Connect(TransCallbackRequest)
	case http.MethodOptions:
		body, err = restTransHandler.Options(TransCallbackRequest)
	case http.MethodTrace:
		body, err = restTransHandler.Trace(TransCallbackRequest)
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

		if err := json.Unmarshal(body, &request.result); nil != err {
			request.result.Fail(err.Error())
		}
	}
	context.JSON(http.StatusOK, request.result)
}
