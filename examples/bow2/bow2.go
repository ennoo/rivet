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
	"github.com/ennoo/rivet/bow"
	"github.com/ennoo/rivet/rivet"
	"github.com/ennoo/rivet/utils/env"
	"strings"
)

func main() {
	rivet.Initialize(true, false, true, false)
	rivet.Bow().Add(
		&bow.RouteService{
			Name:      "test1",
			InURI:     "hello1",
			OutRemote: "http://localhost:8081",
			OutURI:    "rivet/shunt",
		},
		&bow.RouteService{
			Name:      "test2",
			InURI:     "hello2",
			OutRemote: "https://localhost:8092",
			OutURI:    "rivet/shunt",
		},
	)
	rivet.ListenAndServe(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(),
		DefaultPort: "8085",
	}, strings.Join([]string{env.GetEnv(env.GOPath), "/src/github.com/ennoo/rivet/examples/tls/rootCA.crt"}, ""))
}