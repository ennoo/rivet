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
 */

package scheduled

import (
	"github.com/ennoo/rivet/discovery"
	"github.com/ennoo/rivet/server"
	"github.com/ennoo/rivet/utils/slip"
)

// CheckDiscovery 定时检出发现服务中的注册服务列表
func CheckDiscovery(component string) {
	switch component {
	case discovery.ComponentConsul:
		abort := make(chan int, 1)
		go startCheckServicesByConsul(abort)
		a := <-abort
		if execAbortForDiscovery(a) {
			abort = nil
		}
		if abort == nil {
			// 退出前进入自我检查阶段
			CheckServices()
			return
		}
	}
}

// CheckServices 定时检查已存在的服务列表
func CheckServices() {
	abort := make(chan int, 1)
	go startCheckServices(abort)
	a := <-abort
	if execAbortForServices(a) {
		abort = nil
	}
	if abort == nil {
		// todo 退出前解决方案
		return
	}
}

func execAbortForDiscovery(a int) bool {
	switch a {
	case -1: // 启动定时任务出错
		// todo 定时任务出错解决方案
		return true
	case slip.RestResponseError: // 请求对方网络有误
		// todo 请求对方网络有误解决方案
		return true
	case slip.JSONUnmarshalError: // 请求返回数据转JSON失败
		// todo 请求返回数据转JSON失败解决方案
		return true
	}
	return false
}

func execAbortForServices(a int) bool {
	switch a {
	case -1: // 启动定时任务出错
		// todo 定时任务出错解决方案
		return true
	case slip.RestResponseError: // 请求对方网络有误
		// todo 请求对方网络有误解决方案
		return true
	case slip.JSONUnmarshalError: // 请求返回数据转JSON失败
		// todo 请求返回数据转JSON失败解决方案
		return true
	}
	return false
}

// compareAndResetServices 通过比较对象移除原本对象中多余项
func compareAndResetServices(services, servicesCompare *server.Services) {
	for offset := range services.Services {
		have := false
		for position := range servicesCompare.Services {
			if servicesCompare.Services[position].EqualService(services.Services[offset]) {
				have = true
			}
		}
		if !have {
			services.Remove(offset)
		}
	}
}
