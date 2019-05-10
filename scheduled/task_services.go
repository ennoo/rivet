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

package scheduled

import (
	"github.com/ennoo/rivet/discovery/consul"
	"github.com/ennoo/rivet/server"
	"github.com/ennoo/rivet/shunt"
	"github.com/robfig/cron"
)

var abortServices chan int

// startCheckServices 定时检查已存在的服务列表
func startCheckServices(abort chan int) {
	abortServices = abort
	c := cron.New()
	// 每隔5秒执行一次：*/5 * * * * ?
	// 每隔1分钟执行一次：0 */1 * * * ?
	// 每天23点执行一次：0 0 23 * * ?
	// 每天凌晨1点执行一次：0 0 1 * * ?
	// 每月1号凌晨1点执行一次：0 0 1 1 * ?
	// 在26分、29分、33分执行一次：0 26,29,33 * * * ?
	// 每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
	err := c.AddFunc("*/10 * * * * ?", checkServices)
	if nil != err {
		abort <- -1
	} else {
		c.Start()
	}
}

// checkServicesByConsul
//
// 获取本地可负载服务名称列表
//
// 根据本地可负载服务列表遍历发现服务(线上)中是否存在
//
// 如不存在，则继续下一轮遍历
//
// 如存在且列表大于0，遍历线上服务列表并检查线上服务状态是否为可用
//
// 获取本地本地列表 x
//
// 新建空服务列表 y
//
// 如不可用，且 x 中包含此服务，则移除 x 中的服务
//
// 如可用，且 x 中不包含此服务，则新增服务到 x,y 中
//
// 移除 x 中不包含 y 的服务
func checkServices() {
	// 获取本地可负载服务列表
	allWay := shunt.GetShuntInstance().AllWay
	// 根据本地可负载服务列表遍历发现服务(线上)中是否存在
	for serviceName := range allWay {
		agentServiceChecks, slip := consul.ServiceCheck("127.0.0.1:8500", serviceName)
		if nil != slip {
			abortServices <- slip.Code
			return
		}
		// 如不存在，则继续下一轮遍历
		if nil == agentServiceChecks || len(agentServiceChecks) <= 0 {
			continue
		}
		// 获取本地本地列表
		services := server.ServiceGroup()[serviceName]
		if nil == services {
			services = &server.Services{}
			server.ServiceGroup()[serviceName] = services
		}
		// 新建空服务列表
		servicesCompare := server.Services{}
		// 如存在且列表大于0，遍历线上服务列表并检查线上服务状态是否为可用
		checkUpAndLocal(agentServiceChecks, services, &servicesCompare)
		// 移除 x 中不包含 y 的服务
		compareAndResetServices(services, &servicesCompare)
	}
}

// checkUpAndLocalByConsul 如存在且列表大于0，遍历线上服务列表并检查线上服务状态是否为可用
func checkUpAndLocal(agentServiceChecks []*consul.AgentServiceCheck, services, servicesCompare *server.Services) {
	for index := range agentServiceChecks {
		agentServiceCheck := agentServiceChecks[index]
		// 如不可用，且本地列表中包含此服务，则移除本地列表中的服务
		if agentServiceCheck.AggregatedStatus != "passing" {
			for position := range services.Services {
				if services.Services[position].Equal(agentServiceCheck.Service.Address, agentServiceCheck.Service.Port) {
					services.Remove(position)
				}
			}
		} else { // 如可用，且本地列表中不包含此服务，则新增服务到本地列表中
			service := server.Service{
				ID:   agentServiceCheck.Service.ID,
				Host: agentServiceCheck.Service.Address,
				Port: agentServiceCheck.Service.Port,
			}
			have := false
			for position := range services.Services {
				if nil != services.Services && services.Services[position].Equal(service.Host, service.Port) {
					have = true
					break
				}
			}
			if !have {
				services.Add(service)
			}
			servicesCompare.Add(service)
		}
	}
}
