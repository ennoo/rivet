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
	"errors"
	"github.com/ennoo/rivet/server"
)

// 负载均衡接口轮询实现
func init() {
	RegisterBalance("round", &RoundRobinBalance{})
}

type RoundRobinBalance struct {
	curIndex int
}

func (p *RoundRobinBalance) DoBalance(services []*server.Service, key ...string) (add *server.Service, err error) {
	if len(services) == 0 {
		err = errors.New("no instance")
		return
	}

	lens := len(services)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	add = services[p.curIndex]
	p.curIndex++
	return
}
