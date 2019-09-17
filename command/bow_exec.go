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

package command

import (
	"github.com/ennoo/rivet"
	"github.com/ennoo/rivet/bow"
	"github.com/ennoo/rivet/shunt"
	"github.com/ennoo/rivet/trans"
	"github.com/ennoo/rivet/trans/response"
	"github.com/ennoo/rivet/utils/file"
	"strings"
)

func startBow(ymlBow *Bow) error {
	rivet.Initialize(ymlBow.HealthCheck, ymlBow.ServerManager, ymlBow.LoadBalance, false)
	//rivet.Log().Init()
	rivet.UseBow(func(result *response.Result) bool {
		return true
	})
	if ymlBow.DiscoveryInit {
		rivet.UseDiscovery(ymlBow.DiscoveryComponent, ymlBow.DiscoveryURL, "bow", ymlBow.DiscoveryReceiveHost, ymlBow.Port)
	}
	bowConfigPath := ymlBow.ConfigPath
	dataArr, err := file.ReadFileByLine(bowConfigPath)
	if nil != err {
		return err
	}
	data := strings.Join(dataArr, "")
	bytes := []byte(data)

	bow.YamlServices(bytes)

	shunt.YamlLBs(bytes)

	tls := trans.YmlTLS(bytes)
	if ymlBow.OpenTLS {
		rivet.ListenAndServesTLS(&rivet.ListenServe{
			Engine:      rivet.SetupRouter(),
			DefaultPort: "19219",
			CertFile:    tls.TLS.Server.CertFile,
			KeyFile:     tls.TLS.Server.KeyFile,
		}, tls.TLS.Clients)
	} else {
		rivet.ListenAndServes(&rivet.ListenServe{
			Engine:      rivet.SetupRouter(),
			DefaultPort: "19219",
		}, tls.TLS.Clients)
	}
	return nil
}
