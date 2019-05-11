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

package consul

import (
	"github.com/ennoo/rivet/utils/log"
	"go.uber.org/zap"
	"testing"
)

var testLogger *zap.Logger

func logger() {
	testLogger = log.GetLogInstance().New("./logs/consul_test.log", "consul_test")
	testLogger.Info("log consul_test 初始化成功")
}

func TestEnroll(t *testing.T) {
	logger()
	Enroll("127.0.0.1:8500", "ididididid", "rivet", "127.0.0.1", 8080)
}

func TestChecks(t *testing.T) {
	logger()
	Checks()
}

func TestServiceCheck(t *testing.T) {
	logger()
	_, _ = ServiceCheck("operation")
}
