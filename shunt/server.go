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

// 服务器信息
type Address struct {
	host string
	port int
}

func NewAddress(host string, port int) *Address {
	return &Address{
		host: host,
		port: port,
	}
}

func (a *Address) GetHost() string {
	return a.host
}

func (a *Address) GetPort() int {
	return a.port
}
