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
	"fmt"
	"time"
)

func main() {
	tickerStart()
	time.Sleep(time.Second * 10)
}

func tickerStart() {
	ticker := time.NewTicker(time.Millisecond * 350)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("time now = ", time.Now().UnixNano()/1e6)
			}
		}
	}()
}
