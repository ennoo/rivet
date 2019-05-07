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
	"github.com/ennoo/rivet/common/util/log"
	"github.com/ennoo/rivet/rivet"
	"github.com/ennoo/rivet/server"
	"github.com/ennoo/rivet/shunt"
	"github.com/ennoo/rivet/trans/response"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var adds []*server.Service

func main() {
	rivet.Initialize(log.DebugLevel, true, true)
	addAddress()
	rivet.Start(rivet.SetupRouter(testShunt1), "8083")
}

func addAddress() {
	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port, _ := strconv.Atoi(fmt.Sprintf("880%d", i))
		one := server.NewService(host, port)
		adds = append(adds, one)
	}
}

func b() {
	var name = "hash"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for {
		add, err := shunt.DoBalance(name, adds)
		if err != nil {
			fmt.Println("do balance err")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(add)
		time.Sleep(time.Second)
	}
}

func testShunt1(engine *gin.Engine) {
	// 仓库相关路由设置
	vRepo := engine.Group("/rivet3")
	vRepo.GET("/shunt", shunt1)
}

func shunt1(context *gin.Context) {
	rivet.Resp.Do(context, func(result *response.Result) {
		b()
		result.SaySuccess(context, "test2")
	})
}
