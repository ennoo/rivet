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

package env

import (
	"github.com/ennoo/rivet/common/util/string"
	"os"
)

const (
	// ServiceName 当前启动服务名
	ServiceName = "SERVICE_NAME"
	// PortEnv 当前服务启动端口号
	PortEnv = "PORT"
	// ConsulURL 当前服务注册的 consul 地址
	ConsulURL = "CONSUL_URL"
	// Go工作路径
	GOPath = "GOPATH"
)

// GetEnv 获取环境变量 envName 的值
//
// envName 环境变量名称
func GetEnv(envName string) string {
	return os.Getenv(envName)
}

// GetEnvDefault 获取环境变量 envName 的值
//
// envName 环境变量名称
//
// defaultValue 环境变量为空时的默认值
func GetEnvDefault(envName string, defaultValue string) string {
	env := GetEnv(envName)
	if str.IsEmpty(env) {
		return defaultValue
	}
	return env
}
