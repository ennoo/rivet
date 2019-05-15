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

package main

import (
	"fmt"
	"github.com/ennoo/rivet/discovery"
	"github.com/ennoo/rivet/examples/model"
	"github.com/ennoo/rivet/rivet"
	"github.com/ennoo/rivet/trans/request"
	"github.com/ennoo/rivet/trans/response"
	"github.com/ennoo/rivet/utils/ip"
	"github.com/ennoo/rivet/utils/slip"
	"net/http"
)

func main() {
	rivet.Initialize(true, false, false)
	rivet.UseDiscovery(discovery.ComponentConsul, "127.0.0.1:8500", "test", "127.0.0.1", 8081)
	rivet.ListenAndServe(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(testRouter1),
		DefaultPort: "8081",
	})
}

func testRouter1(router *response.Router) {
	// 仓库相关路由设置
	router.Group = router.Engine.Group("/rivet")
	router.GET("/get", shunt1get1)
	router.POST("/post", shunt1post1)
	router.POST("/post2", shunt1post2)
	router.POST("/shunt", shunt1)
}

func shunt1get1(router *response.Router) {
	rivet.Request().Call(router.Context, http.MethodGet, "http://localhost:8082", "rivet/get")
}

func shunt1post1(router *response.Router) {
	fmt.Println("ip = ", ip.Get(router.Context.Request))
	rivet.Request().Callback(router.Context, http.MethodPost, "http://localhost:8082", "rivet/post", func() *response.Result {
		return &response.Result{ResultCode: response.Success, Msg: "降级处理"}
	})
}

func shunt1(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		var test = new(model.Test)
		if err := router.Context.ShouldBindJSON(test); err != nil {
			result.SayFail(router.Context, err.Error())
		}
		test.Name = "trans1"
		router.Context.Writer.Header().Add("trans1Token15", "trans1Test15")
		router.Context.SetCookie("trans1Token16", "trans1Test16", 10, "/", "localhost", false, true)
		result.SaySuccess(router.Context, test)
	})
}

func shunt1post2(router *response.Router) {
	method := http.MethodGet
	remote := "http://127.0.0.1:8500"
	uri := "v1/agent/health/service/name/test"
	_, err := request.SyncPoolGetRequest().RestJSON(method, remote, uri, nil)
	if nil != err {
		slips := err.(*slip.Slip)
		fmt.Println("slips = ", slips.Msg)
	}
}
