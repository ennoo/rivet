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

package cryptos

import "testing"

func TestSha256(t *testing.T) {
	t.Log("------------- mad5 -------------")
	t.Log(MD5("haha"))
	t.Log(MD5("haha"))
	t.Log()
	t.Log("------------- sha1 -------------")
	t.Log(Sha1("haha"))
	t.Log(Sha1("haha"))
	t.Log()
	t.Log("------------- sha224 -------------")
	t.Log(Sha224("haha"))
	t.Log(Sha224("haha"))
	t.Log()
	t.Log("------------- sha256 -------------")
	t.Log(Sha256("haha"))
	t.Log(Sha256("haha"))
	t.Log()
	t.Log("------------- sha384 -------------")
	t.Log(Sha384("haha"))
	t.Log(Sha384("haha"))
	t.Log()
	t.Log("------------- sha512 -------------")
	t.Log(Sha512("haha"))
	t.Log(Sha512("haha"))
	t.Log()
}
