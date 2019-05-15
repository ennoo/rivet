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

package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestBuild(t *testing.T) {
	token, err := Build("1", "rivet", "userMD5", time.Now().Unix(), time.Now().Unix(), time.Now().Unix()+1000)
	if nil == err {
		bo := Check(token)
		fmt.Println("bo = ", bo)
	} else {
		fmt.Println("err = ", err)
	}
}