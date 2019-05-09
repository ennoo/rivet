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
	"github.com/ennoo/rivet/examples/model"
	"github.com/ennoo/rivet/rivet"
	"github.com/ennoo/rivet/trans/response"
	"github.com/gin-gonic/gin"
)

func main() {
	rivet.Initialize(true, false, false)
	rivet.ListenAndServeTLS(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(testRouterTLS2),
		DefaultPort: "8092",
		CertFile:    "/Users/aberic/Documents/tmp/ca/test/server/server.crt",
		KeyFile:     "/Users/aberic/Documents/tmp/ca/test/server/server.key",
	})
}

func testRouterTLS2(engine *gin.Engine) {
	// 仓库相关路由设置
	vRepo := engine.Group("/rivet")
	vRepo.GET("/get", getTLS2)
	vRepo.POST("/post", postTLS2)
	vRepo.POST("/shunt", shuntTLS2)
}

func getTLS2(context *gin.Context) {
	rivet.Response().Do(context, func(result *response.Result) {
		result.SaySuccess(context, "get21")
	})
}

func postTLS2(context *gin.Context) {
	rivet.Response().Do(context, func(result *response.Result) {
		var test = new(model.Test)
		if err := context.ShouldBindJSON(test); err != nil {
			result.SayFail(context, err.Error())
		}
		result.SaySuccess(context, test)
	})
}

func shuntTLS2(context *gin.Context) {
	rivet.Response().Do(context, func(result *response.Result) {
		var test = new(model.Test)
		if err := context.ShouldBindJSON(test); err != nil {
			result.SayFail(context, err.Error())
		}
		test.Name = "trans2"
		result.SaySuccess(context, test)
	})
}
