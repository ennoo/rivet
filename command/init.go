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
	"github.com/ennoo/rivet/utils/file"
	"gopkg.in/yaml.v3"
	"strings"
)

// YmlBow bow object by cmd.yml format
type YmlBow struct {
	Bow Bow `yaml:"bow"`
}

// YmlShunt shunt object by cmd.yml format
type YmlShunt struct {
	Shunt Shunt `yaml:"shunt"`
}

// Bow bow object by cmd.yml format
type Bow struct {
	Port                 int    `yaml:"Port"`                 // 开放端口，便于其它应用访问
	LogPath              string `yaml:"LogPath"`              // 日志文件输出路径
	HealthCheck          bool   `yaml:"HealthCheck"`          // 是否开启健康检查
	ServerManager        bool   `yaml:"ServerManager"`        // 是否启用服务管理功能
	LoadBalance          bool   `yaml:"LoadBalance"`          // 是否启用负载均衡
	OpenTLS              bool   `yaml:"OpenTLS"`              // 是否开启 TLS
	ConfigPath           string `yaml:"ConfigPath"`           // Bow 配置文件路径
	DiscoveryInit        bool   `yaml:"DiscoveryInit"`        // 是否启用发现服务
	DiscoveryComponent   string `yaml:"DiscoveryComponent"`   // 所启用发现服务组件名
	DiscoveryURL         string `yaml:"DiscoveryURL"`         // 发现服务地址
	DiscoveryReceiveHost string `yaml:"DiscoveryReceiveHost"` // 发现服务收到当前注册服务的地址，端口号默认通过 PORT 获取
}

// Shunt shunt object by cmd.yml format
type Shunt struct {
	Port                 int    `yaml:"Port"`                 // 开放端口，便于其它应用访问
	LogPath              string `yaml:"LogPath"`              // 日志文件输出路径
	HealthCheck          bool   `yaml:"HealthCheck"`          // 是否开启健康检查
	ServerManager        bool   `yaml:"ServerManager"`        // 是否启用服务管理功能
	OpenTLS              bool   `yaml:"OpenTLS"`              // 是否开启 TLS
	ConfigPath           string `yaml:"ConfigPath"`           // Bow 配置文件路径
	DiscoveryInit        bool   `yaml:"DiscoveryInit"`        // 是否启用发现服务
	DiscoveryComponent   string `yaml:"DiscoveryComponent"`   // 所启用发现服务组件名
	DiscoveryURL         string `yaml:"DiscoveryURL"`         // 发现服务地址
	DiscoveryReceiveHost string `yaml:"DiscoveryReceiveHost"` // 发现服务收到当前注册服务的地址，端口号默认通过 PORT 获取
}

func formatYml(path string) ([]byte, error) {
	dataArr, err := file.ReadFileByLine(path)
	if nil != err {
		return nil, err
	}
	return []byte(strings.Join(dataArr, "")), nil
}

// YamlBow YML转路由对象
func yamlBow(path string) (*Bow, error) {
	data, err := formatYml(path)
	if nil != err {
		return nil, err
	}
	ymlBow := YmlBow{}
	err = yaml.Unmarshal([]byte(data), &ymlBow)
	if err != nil {
		return nil, err
	}
	return &ymlBow.Bow, nil
}

// yamlShunt YML转路由对象
func yamlShunt(path string) (*Shunt, error) {
	data, err := formatYml(path)
	if nil != err {
		return nil, err
	}
	ymlShunt := YmlShunt{}
	err = yaml.Unmarshal([]byte(data), &ymlShunt)
	if err != nil {
		return nil, err
	}
	return &ymlShunt.Shunt, nil
}
