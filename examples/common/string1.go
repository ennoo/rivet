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
	"strings"
)

func main() {
	formatString1()
	formatString2("http://8b38a7000ce1/actuator/health")
	formatString2("http://127.0.0.1:8555/actuator/health")
}

func formatString1() {
	a := "HTTP GET http://8b38a7000ce1/actuator/health: 200  Output: {\"status\":\"UP\"}"
	as := strings.Split(a, " ")
	res := as[2]
	res = res[0 : len(res)-1]
	fmt.Println("http = ", res)
}

func formatString2(urlIn string) {
	urlTmp := urlIn
	if strings.Contains(urlTmp, "//") {
		urlTmp = strings.Split(urlTmp, "//")[1]
		fmt.Println("url1 = ", urlTmp)
	}
	size := len(strings.Split(urlTmp, "/")[0]) + 1
	urlTmp = urlTmp[size:]
	remote := urlIn[0:(len(urlIn) - len(urlTmp) - 1)]
	uri := urlTmp
	fmt.Println("remote = ", remote)
	fmt.Println("uri = ", uri)
}
