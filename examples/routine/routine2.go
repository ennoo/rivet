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
	"time"
)

func main() {

	c := make(chan int, 2)
	defer close(c)
	//提前将队列放满
	c <- 1
	c <- 2
	fmt.Println("c len = ", len(c))
	fmt.Println("开始尝试执行")
	go cross(c)
	process(c)

}

func process(c chan int) {

	fmt.Println("被限流阻塞")
	c <- 1 //channel已满，将阻塞，直到成功放入channel  **
	fmt.Println("已放行，执行process")

}

func cross(c chan int) {
	a := 1
	b := 5
	for b >= a {
		fmt.Println("阻塞", b, "秒")
		time.Sleep(time.Second)
		b -= 1
	}

	fmt.Println("释放一个通行证")
	fmt.Println("<-c 1 = ", <-c)
	fmt.Println("<-c 2 = ", <-c)
	fmt.Println("<-c 3 = ", <-c)
	fmt.Println("<-c 4 = ", <-c)
	//取出元素，则chan可以继续放入数据，将唤醒**行代码

}
