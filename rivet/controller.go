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
	"github.com/ennoo/rivet/common/discovery"
	"github.com/ennoo/rivet/common/discovery/consul"
	"github.com/ennoo/rivet/common/util/env"
	"github.com/ennoo/rivet/common/util/log"
	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化日志
	log.InitLogger()
	log.Info("controller init")
}

var useDiscovery = false

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
func SetupRouter(healthCheck bool) *gin.Engine {
	engine := gin.Default()
	if healthCheck {
		discovery.Health(engine)
	}
	return engine
}

func Start(engine *gin.Engine, defaultPort string) {
	log.Info("listening port bind")
	err := engine.Run(":" + env.GetEnvDafult(env.PortEnv, defaultPort))
	if nil != err {
		log.Info("exit because {}", err)
	}
}
