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

package server

import (
	"github.com/ennoo/rivet/trans/response"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	Resp     = response.Response{}
	balances = make(map[string]*Services)
)

// Server 服务器管理服务路由
func Server(engine *gin.Engine) {
	// 仓库相关路由设置
	vRepo := engine.Group("/shunt")
	vRepo.GET("/balance/list", listBalance)
	vRepo.DELETE("/balance/rm/:serviceName", rmBalance)
	vRepo.GET("/service/list/:serviceName", listService)
	vRepo.POST("/service/add", addService)
	vRepo.DELETE("/service/rm/:serviceName/:serviceId", rmService)
}

func addService(context *gin.Context) {
	Resp.Do(context, func(result *response.Result) {
		balance := new(Balance)
		if err := context.ShouldBindJSON(balance); err != nil {
			result.SayFail(context, err.Error())
		}
		name := balance.Name
		services := balances[name]
		if nil == services {
			services = &Services{}
			balances[name] = services
		}
		services.Add(balance.Service)
		result.SaySuccess(context, "add service success")
	})
}

func rmService(context *gin.Context) {
	Resp.Do(context, func(result *response.Result) {
		serviceName := context.Param("serviceName")
		serviceId := context.Param("serviceId")
		if nil == balances[serviceName] {
			panic(response.ExpNotExist.Fit(strings.Join([]string{"service", serviceName}, " ")))
		}
		have := false
		service := balances[serviceName]
		services := service.Services
		for i := 0; i < len(services); i++ {
			if serviceId == services[i].Id {
				have = true
				service.Remove(i)
			}
		}
		if have {
			result.SaySuccess(context, strings.Join([]string{"remove service", serviceName, "id =", serviceId, "success"}, " "))
		} else {
			panic(response.ExpNotExist.Fit(strings.Join([]string{"service", serviceName, "id =", serviceId}, " ")))
		}
	})
}

func listService(context *gin.Context) {
	Resp.Do(context, func(result *response.Result) {
		serviceName := context.Param("serviceName")
		if nil != balances[serviceName] {
			result.SaySuccess(context, balances[serviceName].Services)
		} else {
			result.SaySuccess(context, []Service{})
		}
	})
}

func rmBalance(context *gin.Context) {
	Resp.Do(context, func(result *response.Result) {
		serviceName := context.Param("serviceName")
		if nil == balances[serviceName] {
			panic(response.ExpNotExist.Fit(strings.Join([]string{"service", serviceName}, " ")))
		}
		delete(balances, serviceName)
		result.SaySuccess(context, strings.Join([]string{"remove ", serviceName, " balance success"}, ""))
	})
}

func listBalance(context *gin.Context) {
	Resp.Do(context, func(result *response.Result) {
		shunts := make([]string, len(balances))
		index := 0
		for k := range balances {
			shunts[index] = k
			index++
		}
		result.SaySuccess(context, shunts)
	})
}
