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
	"github.com/ennoo/rivet/examples/model"
	"github.com/ennoo/rivet/trans/response"
	"github.com/ennoo/rivet/utils/env"
	"net/http"
	"strings"
	"time"
)

func main() {
	rivet.Initialize(true, false, false)
	rivet.ListenAndServeTLS(&rivet.ListenServe{
		Engine:         rivet.SetupRouter(testRouterTLS1),
		DefaultPort:    "8091",
		ConnectTimeout: 3 * time.Second,
		KeepAlive:      30 * time.Second,
		CertFile:       strings.Join([]string{env.GetEnv(env.GOPath), "/src/github.com/ennoo/rivet/examples/tls/server/server.crt"}, ""),
		KeyFile:        strings.Join([]string{env.GetEnv(env.GOPath), "/src/github.com/ennoo/rivet/examples/tls/server/server.key"}, ""),
	}, strings.Join([]string{env.GetEnv(env.GOPath), "/src/github.com/ennoo/rivet/examples/tls/rootCA.crt"}, ""))
}

func testRouterTLS1(router *response.Router) {
	// 仓库相关路由设置
	router.Group = router.Engine.Group("/rivet")
	router.GET("/get", getTLS1)
	router.POST("/post", postTLS1)
	router.POST("/shunt", shuntTLS1)
}

func getTLS1(router *response.Router) {
	rivet.Request().Call(router.Context, http.MethodGet, "https://localhost:8092", "rivet/get")
}

func postTLS1(router *response.Router) {
	rivet.Request().Callback(router.Context, http.MethodPost, "https://localhost:8092", "rivet/post", func() *response.Result {
		return &response.Result{ResultCode: response.Success, Msg: "降级处理"}
	})
}

func shuntTLS1(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		var test = new(model.Test)
		if err := router.Context.ShouldBindJSON(test); err != nil {
			result.SayFail(router.Context, err.Error())
		}
		test.Name = "trans1"
		result.SaySuccess(router.Context, test)
	})
}
