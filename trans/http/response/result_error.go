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
	"net/http"
)

var (
	// 通用模块
	CommonModel = "02"
	// 仓库管理模块
	DockerModel = "03"

	// Rest调用模块
	RestModel = "04"

	SystemCommonException = NewRespErr("0", "System exception")

	ParamException   = NewRespErr(CommonModel+"00001", "Param exception")
	ServiceException = NewRespErr(CommonModel+"00002", "Service exception,try again later or contact the administrator")

	//rest调用模块
	RestResponseErr = NewRespErr(RestModel+"00001", "接口调用失败，返回错误信息")

	//	Docker模块
	SwarmServiceNameNotExist = NewRespErr(DockerModel+"00001", "Service name [%s] not exist")

	// service的依赖链可能出现了环路
	SwarmServiceDependsonCircle = NewRespErr(DockerModel+"00002", "Service depends on may be referenced with a circle! ")
)

type RespError struct {
	ErrorCode string
	ErrorMsg  string

	// http的错误编码
	HttpStatusCode int
}

func (resErr *RespError) Error() string {
	return fmt.Sprintf("%s,error code is:%s", resErr.ErrorMsg, resErr.ErrorCode)
}

// 添加参数
func (resErr *RespError) FormatMsg(args ...interface{}) *RespError {
	newMsg := fmt.Sprintf(resErr.ErrorMsg, args...)
	resErr.ErrorMsg = newMsg
	return resErr
}

func NewRespErr(errCode string, errMsg string) *RespError {
	return &RespError{
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
		// 默认正常
		HttpStatusCode: http.StatusOK,
	}
}
