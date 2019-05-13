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
	"github.com/ennoo/rivet/rivet"
	"github.com/ennoo/rivet/trans/response"
	"github.com/ennoo/rivet/utils/env"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	rivet.Initialize(false, true, false)
	rivet.UseBow(func(context *gin.Context, result *response.Result) bool {
		result.Fail("test fail")
		return false
	})
	rivet.Bow().AddService("test1", "hello1", "http://localhost:8081", "rivet/shunt")
	rivet.Bow().AddService("test2", "hello2", "https://localhost:8092", "rivet/shunt")
	rivet.ListenAndServe(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(),
		DefaultPort: "8084",
	}, strings.Join([]string{env.GetEnv(env.GOPath), "/src/github.com/ennoo/rivet/examples/tls/rootCA.crt"}, ""))
}
