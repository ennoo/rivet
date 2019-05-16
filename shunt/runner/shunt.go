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

package main

import (
	"github.com/ennoo/rivet"
	"github.com/ennoo/rivet/discovery"
	"github.com/ennoo/rivet/shunt"
	"github.com/ennoo/rivet/trans/response"
	"github.com/ennoo/rivet/utils/env"
	"github.com/ennoo/rivet/utils/log"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strings"
)

func main() {
	rivet.Initialize(true, true, true)
	rivet.Log().Conf(&log.Config{
		FilePath:    strings.Join([]string{"./logs/rivet.log"}, ""),
		Level:       zapcore.DebugLevel,
		MaxSize:     128,
		MaxBackups:  30,
		MaxAge:      30,
		Compress:    true,
		ServiceName: env.GetEnvDefault("SERVICE_NAME", "shunt1"),
	})
	rivet.UseDiscovery(discovery.ComponentConsul, "127.0.0.1:8500", "shunt", "127.0.0.1", 8083)
	rivet.Shunt().Register("test", shunt.Round)
	rivet.Shunt().Register("test1", shunt.Random)
	rivet.Shunt().Register("test2", shunt.Hash)
	//addAddress()
	rivet.ListenAndServe(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(testShunt1),
		DefaultPort: "19219",
	})
}

func testShunt1(router *response.Router) {
	// 仓库相关路由设置
	router.Group = router.Engine.Group("/rivet")
	router.GET("/shunt/:serviceName", shunt3)
	router.POST("/shunt", shunt4)
}

func shunt3(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		serviceName := router.Context.Param("serviceName")
		rivet.Shunt().Register(serviceName, shunt.Round)
		result.SaySuccess(router.Context, "test2")
	})
}

func shunt4(router *response.Router) {
	rivet.Request().Callback(router.Context, http.MethodPost, "test", "rivet/shunt", func() *response.Result {
		return &response.Result{ResultCode: response.Success, Msg: "降级处理"}
	})
}
