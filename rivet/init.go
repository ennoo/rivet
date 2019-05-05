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
	"github.com/gin-gonic/gin"
)

var useDiscovery = false
var hc = false
var st = false

// rivet 初始化方法，必须最先调用
//
// logLevel：日志等级，参考 github/rivet/common/util/log/log.go，一般调试用 DebugLevel，生产用 InfoLevel
//
// healthCheck：是否开启健康检查。开启后为 Get 请求，路径为 /health/check
func Initialize(logLevel string, healthCheck bool, useShunt bool) {
	// 初始化日志
	log.InitLoggerWithLevel(logLevel)
	log.Info("controller init")
	hc = healthCheck
	st = useShunt
}

// 启用指定的发现服务
func UseDiscovery(component string, url string, serviceName string) {
	switch component {
	case discovery.ComponentConsul:
		if !useDiscovery {
			log.Info("use discovery service {}", discovery.ComponentConsul)
			useDiscovery = true
			go consul.Enroll(url, serviceName)
		}
	}
}

// setupRouter设置路由器相关选项
func SetupRouter(routes ...func(*gin.Engine)) *gin.Engine {
	engine := gin.Default()
	if hc {
		Health(engine)
	}
	if st {
		Shunt(engine)
	}
	for _, route := range routes {
		route(engine)
	}
	return engine
}

func Start(engine *gin.Engine, defaultPort string) {
	log.Info("listening port bind")
	err := engine.Run(":" + env.GetEnvDefault(env.PortEnv, defaultPort))
	if nil != err {
		log.Info("exit because {}", err)
	}
}
