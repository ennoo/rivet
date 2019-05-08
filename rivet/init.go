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
	"github.com/ennoo/rivet/common/util/env"
	"github.com/ennoo/rivet/common/util/log"
	"github.com/ennoo/rivet/discovery"
	"github.com/ennoo/rivet/discovery/consul"
	"github.com/ennoo/rivet/server"
	"github.com/ennoo/rivet/trans/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
	"strings"
)

var useDiscovery = false
var hc = false
var sm = false

// Initialize rivet 初始化方法，必须最先调用
//
// healthCheck：是否开启健康检查。开启后为 Get 请求，路径为 /health/check
//
// serverManager：是否开启外界服务管理功能
//
// loadBalance：是否开启负载均衡
func Initialize(healthCheck bool, serverManager bool, loadBalance bool) {
	Log().Conf(&log.Config{
		FilePath:    strings.Join([]string{"./logs/rivet.log"}, ""),
		Level:       zapcore.DebugLevel,
		MaxSize:     128,
		MaxBackups:  30,
		MaxAge:      30,
		Compress:    true,
		ServiceName: serviceName,
	})
	hc = healthCheck
	sm = serverManager
	request.LB = loadBalance
}

// UseDiscovery 启用指定的发现服务
func UseDiscovery(component string, url string, serviceName string) {
	switch component {
	case discovery.ComponentConsul:
		if !useDiscovery {
			log.Rivet.Info("use discovery service {}" + discovery.ComponentConsul)
			useDiscovery = true
			go consul.Enroll(url, serviceName)
		}
	}
}

// SetupRouter 设置路由器相关选项
func SetupRouter(routes ...func(*gin.Engine)) *gin.Engine {
	engine := gin.Default()
	if hc {
		Health(engine)
	}
	if sm {
		server.Server(engine)
	}
	for _, route := range routes {
		route(engine)
	}
	return engine
}

// Start 开始启用 rivet
func Start(engine *gin.Engine, defaultPort string) {
	log.Rivet.Info("listening port bind")
	err := engine.Run(":" + env.GetEnvDefault(env.PortEnv, defaultPort))
	if nil != err {
		log.Rivet.Info("exit because {}" + err.Error())
	}
}
