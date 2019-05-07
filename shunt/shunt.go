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

package shunt

import (
	"fmt"
	"github.com/ennoo/rivet/server"
)

// Way 负载均衡方式接口
type Way interface {

	// Run 负载均衡算法
	Run(string) (*server.Service, error)
}

// Shunt 负载入口对象
type Shunt struct {
	allWay map[string]Way
}

var shunt = Shunt{allWay: make(map[string]Way)}

// Register 注册新的负载方式
func (s *Shunt) Register(serviceName string, way Way) {
	shunt.allWay[serviceName] = way
}

// RunShunt 开启负载
func RunShunt(serviceName string) (*server.Service, error) {
	way, ok := shunt.allWay[serviceName]
	if !ok {
		err := fmt.Errorf("not fount %s", serviceName)
		fmt.Println("not found ", serviceName)
		return nil, err
	}
	service, err := way.Run(serviceName)
	if err != nil {
		err = fmt.Errorf(" %s erros", serviceName)
		return nil, err
	}
	return service, nil
}
