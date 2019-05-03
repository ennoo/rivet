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
	"github.com/ennoo/rivet/common/util/log"
	"github.com/ennoo/rivet/dolphin"
	"github.com/ennoo/rivet/dolphin/http/response"
	"github.com/ennoo/rivet/rivet"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	rivet.Initialize(log.DebugLevel, true)
	rivet.Start(rivet.SetupRouter(testRouter1), "8081")
}

func testRouter1(engine *gin.Engine) {
	// 仓库相关路由设置
	vRepo := engine.Group("/rivet1")
	vRepo.GET("/get", get1)
	vRepo.POST("/post", post1)
}

func get1(context *gin.Context) {
	dolphin.Trans(context, http.MethodGet, "http://localhost:8082", "rivet2/get", nil)
}

func post1(context *gin.Context) {
	dolphin.Trans(context, http.MethodPost, "http://localhost:8082", "rivet2/post", func() *response.Result {
		return &response.Result{ResultCode: response.Success, Msg: "降级处理"}
	})
}
