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

package shunt

import (
	"fmt"
	"github.com/ennoo/rivet/server"
	"github.com/ennoo/rivet/utils/log"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"sync"
)

const (
	// Random 负载均衡 random 策略
	Random = iota
	// Round 负载均衡 round 策略
	Round
	// Hash 负载均衡 hash 策略
	Hash
)

var (
	instance *Shunt
	once     sync.Once
)

// GetShuntInstance 获取负载管理对象 Shunt 单例
func GetShuntInstance() *Shunt {
	once.Do(func() {
		instance = &Shunt{AllWay: make(map[string]int)}
	})
	return instance
}

// Shunt 负载入口对象
type Shunt struct {
	AllWay map[string]int
}

// LBs 负载对象数组
type LBs struct {
	LBs []*LB `yaml:"shunt"`
}

// LB 负载均衡配置对象
type LB struct {
	Name     string `yaml:"Name"`     // 服务名称
	Register int    `yaml:"Register"` // 负载均衡算法，1：随机；2：轮询；3：Hash一致性
}

// YamlLBs YML转负载对象数组
func YamlLBs(data []byte) []*LB {
	lbs := LBs{}
	err := yaml.Unmarshal([]byte(data), &lbs)
	if err != nil {
		log.Bow.Error("cannot unmarshal data: " + err.Error())
		return nil
	}
	return lbs.LBs
}

// Register 注册新的负载方式
func (s *Shunt) Register(serviceName string, way int) {
	switch way {
	case Round:
		if nil == roundRobinBalances {
			roundRobinBalances = make(map[string]*RoundRobinBalance)
		}
		roundRobinBalances[serviceName] = &RoundRobinBalance{
			serviceName: serviceName,
			rrbCh:       generaCount(),
		}
	}
	instance.AllWay[serviceName] = way
}

// RunShunt 开启负载
func RunShunt(serviceName string) (service *server.Service, err error) {
	way, ok := instance.AllWay[serviceName]
	if !ok {
		err := fmt.Errorf("service not fount")
		fmt.Println("not found ", serviceName)
		log.Shunt.Error(err.Error(), zap.String("serviceName", serviceName))
		return nil, err
	}
	switch way {
	case Random:
		service, err = RunRandom(serviceName)
	case Round:
		service, err = RunRound(serviceName)
	case Hash:
		service, err = RunHash(serviceName)
	}
	if err != nil {
		err = fmt.Errorf(" %s erros", serviceName)
		return
	}
	return
}
