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

// Shunt 分流器负载方式对象
type Shunt struct {
	allBalanceWay map[string]BalanceWay
}

var shunt = Shunt{
	allBalanceWay: make(map[string]BalanceWay),
}

func (p *Shunt) registerBalance(name string, b BalanceWay) {
	p.allBalanceWay[name] = b
}

// RegisterBalance 注册新的负载方式
func RegisterBalance(name string, b BalanceWay) {
	shunt.registerBalance(name, b)
}

// DoBalance 开启负载
func DoBalance(name string, services []*server.Service) (add *server.Service, err error) {
	balance, ok := shunt.allBalanceWay[name]
	if !ok {
		err = fmt.Errorf("not fount %s", name)
		fmt.Println("not found ", name)
		return
	}
	add, err = balance.DoBalance(services)
	if err != nil {
		err = fmt.Errorf(" %s erros", name)
		return
	}
	return
}
