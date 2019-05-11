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
	"github.com/gin-gonic/gin"
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

func testRouter1(engine *gin.Engine) {
	// 仓库相关路由设置
	vRepo := engine.Group("/rivet")
	vRepo.GET("/get", shunt1get1)
	vRepo.POST("/post", shunt1post1)
	vRepo.POST("/post2", shunt1post2)
	vRepo.POST("/shunt", shunt1)
}

func shunt1get1(context *gin.Context) {
	rivet.Request().Call(context, http.MethodGet, "http://localhost:8082", "rivet/get")
}

func shunt1post1(context *gin.Context) {
	fmt.Println("ip = ", ip.Get(context.Request))
	rivet.Request().Callback(context, http.MethodPost, "http://localhost:8082", "rivet/post", func() *response.Result {
		return &response.Result{ResultCode: response.Success, Msg: "降级处理"}
	})
}

func shunt1(context *gin.Context) {
	rivet.Response().Do(context, func(result *response.Result) {
		var test = new(model.Test)
		if err := context.ShouldBindJSON(test); err != nil {
			result.SayFail(context, err.Error())
		}
		test.Name = "trans1"
		result.SaySuccess(context, test)
	})
}

func shunt1post2(context *gin.Context) {
	method := http.MethodGet
	remote := "http://127.0.0.1:8500"
	uri := "v1/agent/health/service/name/test"
	_, err := request.SyncPoolGetRequest().RestJSON(method, remote, uri, nil)
	if nil != err {
		slips := err.(*slip.Slip)
		fmt.Println("slips = ", slips.Msg)
	}
}
