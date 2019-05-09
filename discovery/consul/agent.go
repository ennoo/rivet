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

package consul

import (
	"github.com/ennoo/rivet/common/util/env"
	"github.com/ennoo/rivet/common/util/file"
	"github.com/ennoo/rivet/common/util/log"
	str "github.com/ennoo/rivet/common/util/string"
	"github.com/ennoo/rivet/discovery"
	"github.com/ennoo/rivet/trans/request"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func consulRegister(consulURL string, serviceName string, hostname string) {
	if containerID, err := file.ReadFileFirstLine("/etc/hostname"); nil == err && str.IsEmpty(hostname) {
		hostname = containerID
	} else {
		log.Discovery.Info("open /etc/hostname: no such file or directory")
	}
	log.Discovery.Info("serviceID = " + discovery.ServiceID)
	log.Discovery.Info("hostname = " + hostname)

	method := http.MethodPut
	remote := strings.Join([]string{"http://", env.GetEnvDefault(env.ConsulURL, consulURL)}, "")
	uri := "v1/agent/service/register"
	param := Register{
		ID:                discovery.ServiceID,
		Name:              env.GetEnvDefault(env.ServiceName, serviceName),
		Address:           hostname,
		Port:              80,
		EnableTagOverride: false,
		Check: Check{
			DeregisterCriticalServiceAfter: "1m",
			HTTP:                           strings.Join([]string{"http://", hostname, "/health/check"}, ""),
			Interval:                       "10s"},
	}

	body, err := request.SyncPoolGetRequest().RestJSON(method, remote, uri, param)

	if nil != err {
		log.Discovery.Fatal(err.Error(),
			zap.String("url", strings.Join([]string{remote, "/", uri}, "")),
			zap.String("method", method))
	} else {
		log.Discovery.Info(string(body),
			zap.String("url", strings.Join([]string{remote, "/", uri}, "")),
			zap.String("method", method),
		)
	}
}
