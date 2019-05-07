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

package rivet

import (
	"github.com/ennoo/rivet/shunt"
	"github.com/ennoo/rivet/trans/request"
	"github.com/ennoo/rivet/trans/response"
	"sync"
)

var (
	resp = sync.Pool{
		New: func() interface{} {
			return &response.Response{}
		},
	}
	req = sync.Pool{
		New: func() interface{} {
			return &request.Request{}
		},
	}
	// Shunt 负载入口对象
	Shunt = shunt.Shunt{}
)

// Response 提供实例化调用 Do 方法，并内置返回策略
func Response() *response.Response {
	return resp.Get().(*response.Response)
}

// Request 提供实例化调用请求方法，并内置返回策略
func Request() *request.Request {
	return req.Get().(*request.Request)
}
