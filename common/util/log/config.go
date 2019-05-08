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

package log

import "go.uber.org/zap/zapcore"

// Config 日志配置对象
type Config struct {
	FilePath    string        // 日志文件路径
	Level       zapcore.Level // 日志输出级别
	MaxSize     int           // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups  int           // 日志文件最多保存多少个备份
	MaxAge      int           // 文件最多保存多少天
	Compress    bool          // 是否压缩
	ServiceName string        // 日志所属服务名称
}
