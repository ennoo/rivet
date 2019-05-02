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
package str

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
StringIsEmpty 判断字符串是否为空，是则返回true，否则返回false
*/
func IsEmpty(s string) bool {
	return !IsNotEmpty(s)
}

/*
 StringIsNotEmpty 和StringIsEmpty的语义相反
*/
func IsNotEmpty(s string) bool {
	if len(s) == 0 {
		return false
	}
	return true
}

// ParseToStr 将map中的键值对输出成querystring形式
func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}

// 下划线转换，首字母小写变大写，
// 下划线去掉并将下划线后的首字母大写
func Convert(oriString string) string {
	cb := []byte(oriString)
	em := make([]byte, 0, 10)
	b := false
	for i, by := range cb {
		// 首字母如果是小写，则转换成大写
		if i == 0 && (97 <= by && by <= 122) {
			by = by - 32
		} else if by == 95 {
			// 下一个单词要变成大写
			b = true
			continue
		}
		if b {
			if 97 <= by && by <= 122 {
				by = by - 32
			}
			b = false
		}
		em = append(em, by)
	}
	return string(em)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 创建指定长度的随机字符串
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// randSeq16 创建长度为16的随机字符串
func RandSeq16() string {
	return RandSeq(16)
}

// 获取环境变量 envName 的值
//
// envName 环境变量名称
func GetEnv(envName string) string {
	env := os.Getenv(envName)
	if IsEmpty(env) {
		return ""
	}
	return env
}

// 获取环境变量 envName 的值
//
// envName 环境变量名称
//
// defaultValue 环境变量为空时的默认值
func GetEnvDafult(envName string, defaultValue string) string {
	env := os.Getenv(envName)
	if IsEmpty(env) {
		return defaultValue
	}
	return env
}

func Trim(str string) string {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}
