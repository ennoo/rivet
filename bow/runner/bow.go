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
	"github.com/ennoo/rivet"
	"github.com/ennoo/rivet/bow"
	"github.com/ennoo/rivet/trans/response"
	"github.com/ennoo/rivet/utils/env"
	"github.com/ennoo/rivet/utils/log"
	"go.uber.org/zap/zapcore"
	"strings"
)

func main() {
	rivet.Initialize(false, true, false)
	rivet.Log().Conf(&log.Config{
		FilePath:    strings.Join([]string{"./logs/rivet.log"}, ""),
		Level:       zapcore.DebugLevel,
		MaxSize:     128,
		MaxBackups:  30,
		MaxAge:      30,
		Compress:    true,
		ServiceName: env.GetEnvDefault("SERVICE_NAME", "shunt1"),
	})
	rivet.UseBow(func(result *response.Result) bool {
		return true
	})
	rivet.Bow().Add(
		&bow.RouteService{
			Name:      "test1",
			InURI:     "hello1",
			OutRemote: "http://localhost:8081",
			OutURI:    "rivet/shunt",
			Limit: &bow.Limit{
				LimitMillisecond:         int64(3 * 1000),
				LimitCount:               3,
				LimitIntervalMillisecond: 150,
				LimitChan:                make(chan int, 10),
			},
		},
		&bow.RouteService{
			Name:      "test2",
			InURI:     "hello2",
			OutRemote: "https://localhost:8092",
			OutURI:    "rivet/shunt",
			Limit: &bow.Limit{
				LimitMillisecond:         int64(3 * 1000),
				LimitCount:               3,
				LimitIntervalMillisecond: 150,
				LimitChan:                make(chan int, 10),
			},
		},
	)
	rivet.ListenAndServe(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(),
		DefaultPort: "19219",
	}, strings.Join([]string{env.GetEnv(env.GOPath), "/src/github.com/ennoo/rivet/examples/tls/rootCA.crt"}, ""))
}
