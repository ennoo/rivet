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

package bow

import (
	"fmt"
	"testing"
	"time"
)

func TestLimit(t *testing.T) {
	limit := Limit{
		LimitMillisecond:         int64(1 * 1000),
		LimitCount:               3,
		LimitIntervalMillisecond: 150,
		LimitChan:                make(chan int, 10),
	}
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)
	limit.timeInit()
	fmt.Printf("%v\n", limit.Times)
	time.Sleep(1 * time.Second)
	limit.remove()
	fmt.Printf("%v\n", limit.Times)
	time.Sleep(1 * time.Second)
	limit.add(time.Now().UnixNano() / 1e6)
	fmt.Printf("%v\n", limit.Times)
	go limit.limit()
	loop(&limit)
}

func loop(limit *Limit) {
	i := 0
	for i <= 20 {
		fmt.Println("被堵住了 c len = ", len(limit.LimitChan))
		limit.LimitChan <- i
		fmt.Printf("被放行了 %v\n", limit.Times)
		i++
		time.Sleep(100 * time.Millisecond)
	}
}

func TestLimitMap(t *testing.T) {
	serviceName := "test"
	routeService := RouteService{
		Name:      serviceName,
		InURI:     "hello1",
		OutRemote: "http://localhost:8081",
		OutURI:    "rivet/shunt",
		Limit: &Limit{
			LimitMillisecond:         int64(1 * 1000),
			LimitCount:               3,
			LimitIntervalMillisecond: 150,
			LimitChan:                make(chan int, 10),
		},
	}
	go routeService.Limit.limit()
	loopMap(routeService.Limit)
}

func loopMap(limit *Limit) {
	if nil != limit {
		loop(limit)
	}
}
