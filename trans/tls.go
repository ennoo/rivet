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

package trans

import (
	"github.com/ennoo/rivet/utils/log"
	"gopkg.in/yaml.v3"
)

// TLSYml YML TLS 对象
type TLSYml struct {
	TLS TLS `yaml:"tls"`
}

// TLS 请求服务端及客户端文件对象
type TLS struct {
	Server  *Server  `yaml:"server"`
	Clients []string `yaml:"clients"`
}

// Server TLS 请求服务端文件对象
type Server struct {
	CertFile string `yaml:"CertFile"`
	KeyFile  string `yaml:"KeyFile"`
}

// YmlTLS YML 转 TLS 对象
func YmlTLS(data []byte) *TLSYml {
	tlsYml := TLSYml{}
	err := yaml.Unmarshal([]byte(data), &tlsYml)
	if err != nil {
		log.Trans.Error("cannot unmarshal data: " + err.Error())
		return nil
	}
	return &tlsYml
}
