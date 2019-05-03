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

type BalanceMgr struct {
	allBalance map[string]Balance
}

var mgr = BalanceMgr{
	allBalance: make(map[string]Balance),
}

func (p *BalanceMgr) registerBalance(name string, b Balance) {
	p.allBalance[name] = b
}

func RegisterBalance(name string, b Balance) {
	mgr.registerBalance(name, b)
}

func DoBalance(name string, insts []*Address) (inst *Address, err error) {
	balance, ok := mgr.allBalance[name]
	if !ok {
		err = fmt.Errorf("not fount %s", name)
		fmt.Println("not found ", name)
		return
	}
	inst, err = balance.DoBalance(insts)
	if err != nil {
		err = fmt.Errorf(" %s erros", name)
		return
	}
	return
}
