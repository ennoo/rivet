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
	"github.com/ennoo/rivet/dolphin/http/response"
	"github.com/ennoo/rivet/rivet"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := rivet.SetupRouter(true)
	testRouter(engine)
	rivet.Start(engine, "80")
}

func testRouter(engine *gin.Engine) {
	// 仓库相关路由设置
	vRepo := engine.Group("/rivet")
	vRepo.GET("/test", test)
}

func test(engine *gin.Context) {
	response.Do(engine, nil, func() (i interface{}, e error) {
		return "test", nil
	})
}
