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
	"github.com/ennoo/rivet/shunt"
	"github.com/ennoo/rivet/trans"
	"github.com/ennoo/rivet/utils/file"
	"strings"
)

func startShunt(ymlShunt *Shunt) error {
	rivet.Initialize(ymlShunt.HealthCheck, ymlShunt.ServerManager, true, false)
	//rivet.Log().Init()
	if ymlShunt.DiscoveryInit {
		rivet.UseDiscovery(ymlShunt.DiscoveryComponent, ymlShunt.DiscoveryURL, "shunt", ymlShunt.DiscoveryReceiveHost, ymlShunt.Port)
	}
	bowConfigPath := ymlShunt.ConfigPath
	dataArr, err := file.ReadFileByLine(bowConfigPath)
	if nil != err {
		return err
	}
	data := strings.Join(dataArr, "")
	bytes := []byte(data)

	shunt.YamlLBs(bytes)

	tls := trans.YmlTLS(bytes)
	if ymlShunt.OpenTLS {
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
	return nil
}
