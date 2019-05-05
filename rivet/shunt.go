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

package rivet

import (
	"github.com/ennoo/rivet/common"
	"github.com/ennoo/rivet/shunt"
	"github.com/gin-gonic/gin"
	"reflect"
)

var balance = make(map[string]*shunt.Services)

func Shunt(engine *gin.Engine) {
	// 仓库相关路由设置
	vRepo := engine.Group("/shunt")
	vRepo.GET("/balance/list", listBalance)
	vRepo.POST("/service/add", addService)
	vRepo.POST("/service/list", listService)
}

func addService(engine *gin.Context) {
	Resp.Do(engine, new(shunt.Balance), func(value interface{}) (interface{}, error) {
		result := reflect.ValueOf(value).Elem()
		name := result.FieldByName("Name").String()
		services := balance[name]
		if nil == services {
			services = &shunt.Services{}
			balance[name] = services
		}
		service := shunt.Service{
			Host: result.FieldByName("Service").FieldByName("Host").String(),
			Port: int(result.FieldByName("Service").FieldByName("Port").Int())}
		services.Add(service)
		return nil, nil
	})
}

func listService(engine *gin.Context) {
	Resp.Do(engine, new(common.JsonString), func(value interface{}) (interface{}, error) {
		result := reflect.ValueOf(value).Elem()
		return balance[result.FieldByName("Value").String()], nil
	})
}

func listBalance(engine *gin.Context) {
	Resp.Do(engine, nil, func(value interface{}) (interface{}, error) {
		shunts := make([]string, len(balance))
		index := 0
		for k := range balance {
			shunts[index] = k
			index++
		}
		return shunts, nil
	})
}
