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
	"github.com/ennoo/rivet/utils/log"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Commencing coutdown")
	abort1 := make(chan int, 1)
	abort2 := make(chan int, 1)
	var abort3 chan int
	go func() {
		go func() {
			for i := 0; i < 10; i++ {
				time.Sleep(1 * time.Second)
				fmt.Println("go2 = " + strconv.Itoa(i))
			}
			abort2 <- 1
		}()
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("go1 = " + strconv.Itoa(i))
			abort1 <- i
		}
	}()
	timeout := time.After(5 * time.Second) // timeout 是一个计时信道, 如果达到时间了，就会发一个信号出来
	for isTimeout := false; !isTimeout; {
		select {
		case m := <-abort1:
			if m == 3 {
				abort1 = nil
			}
		case <-abort2:
			abort2 = nil
		case <-abort3:

		case <-timeout:
			log.Scheduled.Debug("超时")
			isTimeout = true // 超时
		}
		if abort2 == nil {
			return
		}
	}

}
