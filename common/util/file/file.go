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

package file

import (
	"bufio"
	"io"
	"os"
)

// PathExists 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ReadFileByLine 从文件中逐行读取并返回字符串数组
func ReadFileByLine(filePath string) ([]string, error) {
	fileIn, fileInErr := os.Open(filePath)
	if fileInErr != nil {
		return nil, fileInErr
	}
	defer fileIn.Close()
	finReader := bufio.NewReader(fileIn)
	var fileList []string
	for {
		inputString, readerError := finReader.ReadString('\n')
		//fmt.Println(inputString)
		if readerError == io.EOF {
			break
		}
		fileList = append(fileList, inputString)
	}
	//fmt.Println("fileList",fileList)
	return fileList, nil
}
