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

import "fmt"

type Shunt struct {
	allBalance map[string]Balance
}

var shunt = Shunt{
	allBalance: make(map[string]Balance),
}

func (p *Shunt) registerBalance(name string, b Balance) {
	p.allBalance[name] = b
}

func RegisterBalance(name string, b Balance) {
	shunt.registerBalance(name, b)
}

func DoBalance(name string, services []*Service) (add *Service, err error) {
	balance, ok := shunt.allBalance[name]
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
