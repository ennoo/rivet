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
	"github.com/ennoo/rivet"
	"github.com/ennoo/rivet/discovery"
	"github.com/ennoo/rivet/shunt"
	"github.com/ennoo/rivet/trans/response"
	"github.com/ennoo/rivet/utils/env"
	"strings"
)

func main() {
	rivet.Initialize(true, true, true)
	rivet.UseDiscovery(discovery.ComponentConsul, "127.0.0.1:8500", "bow", "127.0.0.1", 19219)
	rivet.Shunt().Register("test", shunt.Round)
	rivet.UseBow(func(result *response.Result) bool {
		//result.Fail("test fail")
		//return false
		return true
	})
	rivet.Bow().AddService("test", "hello", "test")
	rivet.Bow().AddService("test1", "hello1", "http://localhost:8081")
	rivet.Bow().AddService("test2", "hello2", "https://localhost:8092")
	rivet.ListenAndServe(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(),
		DefaultPort: "19219",
	}, strings.Join([]string{env.GetEnv(env.GOPath), "/src/github.com/ennoo/rivet/examples/tls/rootCA.crt"}, ""))
}
