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
package consul

import (
	"fmt"
	"github.com/ennoo/rivet/common/util/env"
	"github.com/ennoo/rivet/common/util/file"
	"github.com/ennoo/rivet/common/util/log"
	"github.com/ennoo/rivet/common/util/string"
	"github.com/ennoo/rivet/trans/request"
	"github.com/rs/xid"
	"os"
	"strings"
)

var ServiceID = xid.New().String()

// 调用此方法注册 consul
//
// consulUrl：consul 注册地址，包括端口号（优先通过环境变量 CONSUL_URL 获取）
//
// serviceName：注册到 consul 的服务名称（优先通过环境变量 SERVICE_NAME 获取）
func Enroll(consulUrl string, serviceName string) {
	defer func() {
		log.Info("register sul start")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
			os.Exit(0)
		}
	}()
	consulRegister(consulUrl, serviceName)
}

func consulRegister(consulUrl string, serviceName string) {
	hosts, err := file.ReadFileByLine("/etc/hostname")
	if nil != err {
		panic(err)
	}
	log.Info("serviceID = ", ServiceID)
	containerID := str.Trim(hosts[0])
	log.Info("containerID = ", containerID)
	restJsonHandler := request.RestJsonHandler{
		Param: Register{
			ID:                ServiceID,
			Name:              env.GetEnvDafult(env.ServiceName, serviceName),
			Address:           containerID,
			Port:              80,
			EnableTagOverride: false,
			Check: Check{
				DeregisterCriticalServiceAfter: "1m",
				HTTP:                           strings.Join([]string{"http://", containerID, "/health/check"}, ""),
				Interval:                       "10s"}},
		RestHandler: request.RestHandler{
			RemoteServer: env.GetEnvDafult(env.ConsulUrl, consulUrl),
			Uri:          "/v1/agent/service/register",
			Header:       nil,
			Cookies:      nil}}
	body, err := restJsonHandler.Put()
	if nil != err {
		log.Error(err.Error())
	}
	log.Info("register result = ", string(body))
}
