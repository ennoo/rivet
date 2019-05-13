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

package bow

import (
	"fmt"
	"github.com/ennoo/rivet/trans/request"
	"github.com/ennoo/rivet/trans/response"
	"github.com/ennoo/rivet/utils/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

var (
	instance      *Bow
	once          sync.Once
	serviceCount  = 0
	routeServices = make(map[string]*RouteService)
)

// GetBowInstance 获取路由管理对象 Bow 单例
func GetBowInstance() *Bow {
	once.Do(func() {
		instance = &Bow{AllWay: make(map[string]*RouteService)}
	})
	return instance
}

// Bow 路由入口对象
type Bow struct {
	AllWay map[string]*RouteService
}

// RouteService 路由对象
type RouteService struct {
	Name      string
	InURI     string
	OutRemote string
	OutURI    string
}

// Add 新增路由服务数组
func (s *Bow) Add(routeServiceArr ...*RouteService) {
	for index := range routeServiceArr {
		routeService := routeServiceArr[index]
		routeServices[routeService.Name] = routeService
		GetBowInstance().register(routeService)
		serviceCount++
	}
}

// AddService 新增路由服务
func (s *Bow) AddService(serviceName, inURI, outRemote, outURI string) {
	routeServices[serviceName] = &RouteService{
		Name:      serviceName,
		InURI:     inURI,
		OutRemote: outRemote,
		OutURI:    outURI,
	}
	GetBowInstance().register(&RouteService{
		Name:      serviceName,
		InURI:     inURI,
		OutRemote: outRemote,
		OutURI:    outURI,
	})
	serviceCount++
}

// Register 注册新的路由方式
func (s *Bow) register(routeService *RouteService) {
	instance.AllWay[routeService.Name] = routeService
}

// RunBow 开启路由
func RunBow(context *gin.Context, serviceName string, filter func(context *gin.Context, result *response.Result) bool) {
	RunBowCallback(context, serviceName, filter, nil)
}

// RunBowCallback 开启路由并处理降级
func RunBowCallback(context *gin.Context, serviceName string, filter func(context *gin.Context, result *response.Result) bool, f func() *response.Result) {
	routeService, ok := instance.AllWay[serviceName]
	result := response.Result{}
	if !ok {
		err := fmt.Errorf("routeService not fount")
		log.Shunt.Error(err.Error(), zap.String("serviceName", serviceName))
		result.Fail(err.Error())
		context.JSON(http.StatusOK, result)
		return
	}
	if !filter(context, &result) {
		context.JSON(http.StatusOK, result)
		return
	}
	if nil == f {
		request.SyncPoolGetRequest().Call(context, context.Request.Method, routeService.OutRemote, routeService.OutURI)
	} else {
		request.SyncPoolGetRequest().Callback(context, context.Request.Method, routeService.OutRemote, routeService.OutURI, f)
	}
}
