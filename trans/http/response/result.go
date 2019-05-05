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
	if nil == callback || str.IsEmpty(callback().ResultCode) {
		log.Info("放弃降级或降级策略有误")
		result.ResultCode = Fail
		result.Msg = err.Error()
	} else {
		log.Info("降级回调")
		result.reSet(callback())
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
