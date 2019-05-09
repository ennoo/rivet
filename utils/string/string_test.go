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

package str

import (
	"fmt"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	fmt.Println("haha empty =", IsEmpty("haha"))
	fmt.Println("'' empty =", IsEmpty(""))
}

func TestIsNotEmpty(t *testing.T) {
	fmt.Println("haha empty =", IsNotEmpty("haha"))
	fmt.Println("'' empty =", IsNotEmpty(""))
}

func TestConvert(t *testing.T) {
	fmt.Println("uu_xx_aa =", Convert("uu_xx_aa"))
}

func TestRandSeq(t *testing.T) {
	fmt.Println("13 =", RandSeq(13))
	fmt.Println("23 =", RandSeq(23))
	fmt.Println("33 =", RandSeq(33))
}

func TestRandSeq16(t *testing.T) {
	fmt.Println("RandSeq16 =", RandSeq16())
}

func TestTrim(t *testing.T) {
	s := "kjsdhfj ajsd\nksjhdka sjkh"
	fmt.Println(s, "=", Trim(s))
}
