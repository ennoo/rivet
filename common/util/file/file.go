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
	"go.uber.org/zap"
	"io"
	"os"
)

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

// 创建指定目录
// paths代表待创建的目录的路径,可以同时创建多个路径。
// ignoreExist代表是否忽略已存在的目录，如果为true，则代表忽略。即如果目录已存在也不会报错。
// 如果该值为false，则如果指定目录已存在，将会报错。
func DirectoryCreate(paths ...string) (err error) {
	//如果报错，则停止后续的动作
	for _, dirName := range paths {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			panic(err)
			return err
		}
		zap.S().Debug("Create path success :" + dirName)
	}
	return err
}

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
