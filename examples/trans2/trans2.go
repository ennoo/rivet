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
	"github.com/ennoo/rivet"
	"github.com/ennoo/rivet/discovery"
	"github.com/ennoo/rivet/examples/model"
	"github.com/ennoo/rivet/trans/response"
)

func main() {
	rivet.Initialize(true, false, false)
	rivet.UseDiscovery(discovery.ComponentConsul, "127.0.0.1:8500", "test", "127.0.0.1", 8082)
	rivet.ListenAndServe(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(testRouter2),
		DefaultPort: "8082",
	})
}

func testRouter2(router *response.Router) {
	// 仓库相关路由设置
	router.Group = router.Engine.Group("/rivet")
	router.GET("/get", get2)
	router.POST("/post", post2)
	router.POST("/shunt", shunt2)
}

func get2(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		result.SaySuccess(router.Context, "get21")
	})
}

func post2(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		var test = new(model.Test)
		if err := router.Context.ShouldBindJSON(test); err != nil {
			result.SayFail(router.Context, err.Error())
		}
		result.SaySuccess(router.Context, test)
	})
}

func shunt2(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		var test = new(model.Test)
		if err := router.Context.ShouldBindJSON(test); err != nil {
			result.SayFail(router.Context, err.Error())
		}
		test.Name = "trans2"
		result.SaySuccess(router.Context, test)
	})
}
