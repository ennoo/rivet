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

// Package consul 支持包
package consul

import (
	"fmt"
	"github.com/ennoo/rivet/common/util/log"
	"os"
)

// Enroll 调用此方法注册 consul
//
// consulUrl：consul 注册地址，包括端口号（优先通过环境变量 CONSUL_URL 获取）
//
// serviceName：注册到 consul 的服务名称（优先通过环境变量 SERVICE_NAME 获取）
//
// hostname：注册到 consul 的服务地址（如果为空，则尝试通过 /etc/hostname 获取）
func Enroll(consulURL string, serviceName string, hostname string) {
	defer func() {
		log.Discovery.Info("register sul start")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
			os.Exit(0)
		}
	}()
	consulRegister(consulURL, serviceName, hostname)
}
