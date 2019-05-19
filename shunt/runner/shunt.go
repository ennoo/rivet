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
	"github.com/ennoo/rivet/discovery"
	"github.com/ennoo/rivet/shunt"
	"github.com/ennoo/rivet/trans"
	"github.com/ennoo/rivet/utils/env"
	"github.com/ennoo/rivet/utils/file"
	"github.com/ennoo/rivet/utils/log"
	"go.uber.org/zap"
	"strings"
)

func main() {
	rivet.Initialize(env.GetEnvBoolDefault(env.HealthCheck, false),
		env.GetEnvBoolDefault(env.ServerManager, false), true)
	rivet.Log().Init()
	if env.GetEnvBoolDefault(env.DiscoveryInit, false) {
		rivet.UseDiscovery(discovery.ComponentConsul, "127.0.0.1:8500", "shunt", "127.0.0.1", 8083)
	}

	bowConfigPath := env.GetEnv(env.ConfigPath)
	dataArr, err := file.ReadFileByLine(bowConfigPath)
	if nil != err {
		log.Shunt.Panic("load bow config yml failed", zap.String("BOW_CONFIG_PATH", bowConfigPath), zap.Error(err))
	}
	data := strings.Join(dataArr, "")
	log.Shunt.Debug("yml string", zap.String("data", data))
	bytes := []byte(data)

	shunt.YamlLBs(bytes)

	tls := trans.YmlTLS(bytes)
	if env.GetEnvBool(env.OpenTLS) {
		rivet.ListenAndServesTLS(&rivet.ListenServe{
			Engine:      rivet.SetupRouter(),
			DefaultPort: "19877",
			CertFile:    tls.TLS.Server.CertFile,
			KeyFile:     tls.TLS.Server.KeyFile,
		}, tls.TLS.Clients)
	} else {
		rivet.ListenAndServes(&rivet.ListenServe{
			Engine:      rivet.SetupRouter(),
			DefaultPort: "19877",
		}, tls.TLS.Clients)
	}
}
