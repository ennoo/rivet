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
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

//
func Do(c *gin.Context, obj interface{}, f func() (interface{}, error)) {
	res := Result{}
	defer catchErr(c, &res)
	//var deployModel = new(model.DeployModel)
	if nil != obj {
		if err := c.ShouldBindJSON(obj); err != nil {
			res.failErr(err)
			c.JSON(http.StatusOK, &res)
			return
		}
	}
	result, err := f()
	exec(&res, c, err, result)
}

func DoNoData(c *gin.Context, obj interface{}, f func() error) {
	res := Result{}
	defer catchErr(c, &res)
	if nil != obj {
		if err := c.ShouldBindJSON(obj); err != nil {
			res.failErr(err)
			c.JSON(http.StatusOK, &res)
			return
		}
	}
	execWithNoData(&res, c, f())
}

//Success 默认成功返回
func (result *Result) success(obj interface{}) {
	result.Msg = "Success!"
	result.ResultCode = Success
	result.Data = obj
}

// Fail 方法主要提供返回错误的json数据
func (result *Result) fail(msg string) {
	result.Msg = msg
	result.ResultCode = Fail
}

// FailErr 携带error信息,如果是respError，则
// 必然存在errorCode和msg，因此进行赋值。否则不赋值
func (result *Result) failErr(err error) {
	switch vtype := err.(type) {
	case *RespError:
		result.Msg = vtype.ErrorMsg
		result.ResultCode = vtype.ErrorCode
	default:
		result.ResultCode = Fail
		zap.S().Error(err)
		//result.Msg = ServiceException.ErrorMsg
		result.Msg = err.Error()
	}

}

func execWithNoData(res *Result, c *gin.Context, err error) {
	exec(res, c, err, nil)
}

func exec(res *Result, c *gin.Context, err error, resObj interface{}) {
	if err != nil {
		res.fail(err.Error())
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, &res)
		return
	}
	if nil != resObj {
		res.success(resObj)
	}
	c.JSON(http.StatusOK, &res)
}

//捕获所有异常信息并放入json到context，便于controller直接调用
func catchErr(c *gin.Context, res *Result) {
	if r := recover(); r != nil {
		//fmt.Printf("捕获到的错误：%s\n", r)
		res.fail(fmt.Sprintf("An error occurred:%v \n", r))
		zap.S().Error(r)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
}
