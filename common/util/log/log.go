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

package log

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

const (
	// DebugLevel 日志级别为 debug
	DebugLevel = "debug"
	// InfoLevel 日志级别为 info
	InfoLevel = "info"
)

// Initialize 初始化日志组件并指定日志级别
func Initialize(level string) {
	var err error
	if level == DebugLevel {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	logger.Debug("Logger start success")
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	zap.S().Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	zap.S().Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	zap.S().Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	zap.S().Error(args)
}
