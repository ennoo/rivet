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
	"fmt"
	"github.com/ennoo/rivet/common/util/log"
	"github.com/ennoo/rivet/common/util/string"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success = "200"
	Fail    = "9999"
)

type Result struct {
	ResultCode string `json:"code"`
	Msg        string `json:"msg"`
	// 数据接口
	Data interface{} `json:"data"`
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
func Do(context *gin.Context, obj interface{}, objBlock func(value interface{}) (interface{}, error)) {
	res := Result{}
	defer catchErr(context, &res)
	//var deployModel = new(model.DeployModel)
	if nil != obj {
		if err := context.ShouldBindJSON(obj); err != nil {
			res.FailErr(err)
			context.JSON(http.StatusOK, &res)
			return
		}
	}
	result, err := objBlock(obj)
	exec(&res, context, err, result)
}

//Success 默认成功返回
func (result *Result) Success(obj interface{}) {
	result.Msg = "Success!"
	result.ResultCode = Success
	result.Data = obj
}

// Fail 方法主要提供返回错误的json数据
func (result *Result) Fail(msg string) {
	result.Msg = msg
	result.ResultCode = Fail
}

// 返回结果对象介入降级操作方法
func (result *Result) Callback(callback func() *Result, err error) {
	resultCallBack := callback()
	if str.IsNotEmpty(resultCallBack.ResultCode) {
		log.Info("降级回调")
		result.reSet(resultCallBack)
	} else {
		log.Info("放弃降级或降级策略有误")
		result.ResultCode = Fail
		result.Msg = err.Error()
	}
}

func (result *Result) reSet(res *Result) {
	result.ResultCode = res.ResultCode
	result.Data = res.Data
	result.Msg = res.Msg
}

// FailErr 携带error信息,如果是respError，则
// 必然存在errorCode和msg，因此进行赋值。否则不赋值
func (result *Result) FailErr(err error) {
	switch vtype := err.(type) {
	case *RespError:
		result.Msg = vtype.ErrorMsg
		result.ResultCode = vtype.ErrorCode
	default:
		result.ResultCode = Fail
		log.Error(err)
		//result.Msg = ServiceException.ErrorMsg
		result.Msg = err.Error()
	}

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

//捕获所有异常信息并放入json到context，便于controller直接调用
func catchErr(context *gin.Context, res *Result) {
	if r := recover(); r != nil {
		//fmt.Printf("捕获到的错误：%s\n", r)
		res.Fail(fmt.Sprintf("An error occurred:%v \n", r))
		log.Error(r)
		context.JSON(http.StatusInternalServerError, res)
		return
	}
}
