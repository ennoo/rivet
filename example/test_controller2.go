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
	"github.com/ennoo/rivet/example/model"
	"github.com/ennoo/rivet/rivet"
	"github.com/ennoo/rivet/trans/http/response"
	"github.com/gin-gonic/gin"
)

func main() {
	rivet.Initialize(log.DebugLevel, true)
	rivet.Start(rivet.SetupRouter(testRouter2), "8082")
}

func testRouter2(engine *gin.Engine) {
	// 仓库相关路由设置
	vRepo := engine.Group("/rivet2")
	vRepo.GET("/get", get2)
	vRepo.POST("/post", post2)
}

func get2(context *gin.Context) {
	response.Do(context, nil, func(value interface{}) (interface{}, error) {
		return "test2", nil
	})
}

func post2(context *gin.Context) {
	response.Do(context, new(model.Test), func(value interface{}) (interface{}, error) {
		return value, nil
	})
}
