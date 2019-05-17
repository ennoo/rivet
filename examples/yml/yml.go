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

package main

import (
	"fmt"
	"github.com/ennoo/rivet/utils/file"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// T Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string `yaml:"a"`
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:"d,flow"`
	}
}

func main() {
	t := T{}

	datas, _ := file.ReadFileByLine("/Users/admin/Documents/code/git/go/src/github.com/ennoo/rivet/examples/yml/test.yml")
	data := strings.Join(datas, "")

	//data := file.ReadFile("/Users/admin/Documents/code/git/go/src/github.com/ennoo/rivet/examples/yml/test.yml")
	fmt.Println("data : ", data)
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
