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
package response

import (
	"github.com/ennoo/rivet/common/util/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	result Result
}

// 处理 request 请求
//
// context：请求上下文
//
// obj：请求中 body 中内容期望转换的对象并做空实例化，如 new(Type)
//
// objBlock：obj 对象的回调方法，最终调用 Do 函数的方法会接收到返回值
//
// objBlock interface{}：obj 对象的回调方法所返回的最终交由 response 输出的对象
//
// objBlock error：obj 对象的回调方法所返回的错误对象
//
// 如未出现err，且无可描述返回内容，则返回值可为 (nil, nil)
func (response *Response) Do(context *gin.Context, obj interface{}, objBlock func(value interface{}) (interface{}, error)) {
	defer catchErr(context, &response.result)
	//var deployModel = new(model.DeployModel)
	if nil != obj {
		if err := context.ShouldBindJSON(obj); err != nil {
			response.result.FailErr(err)
			context.JSON(http.StatusOK, &response.result)
			return
		}
	}
	result, err := objBlock(obj)
	exec(&response.result, context, err, result)
}

// 处理结果并给出对应返回策略
func exec(res *Result, context *gin.Context, err error, resObj interface{}) {
	if err != nil {
		res.Fail(err.Error())
		log.Error(err)
		context.JSON(http.StatusInternalServerError, &res)
		return
	}
	if nil != resObj {
		res.Success(resObj)
	}
	context.JSON(http.StatusOK, &res)
}
