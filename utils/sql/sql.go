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

package sql

import (
	"github.com/ennoo/rivet/utils/env"
	"github.com/ennoo/rivet/utils/log"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron"
	"strings"
	"time"
)

var (
	// 数据库任务入口
	db *gorm.DB
	// dbURL 数据库 URL
	dbURLSelf string
	// dbUser 数据库用户名
	dbUserSelf string
	// dbPass 数据库用户密码
	dbPassSelf string
	// dbName 数据库名称
	dbNameSelf string
)

// Connect 链接数据库服务
func Connect(dbURL, dbUser, dbPass, dbName string) error {
	if nil == db {
		dbURLSelf = env.GetEnvDefault(env.DBURL, dbURL)
		dbUserSelf = env.GetEnvDefault(env.DBUser, dbUser)
		dbPassSelf = env.GetEnvDefault(env.DBPass, dbPass)
		dbNameSelf = env.GetEnvDefault(env.DBName, dbName)
		log.SQL.Warn("init DB Manager")
		dbValue := strings.Join([]string{dbUserSelf, ":", dbPassSelf, "@tcp(", dbURLSelf, ")/", dbNameSelf,
			"?charset=utf8&parseTime=True&loc=Local"}, "")
		var err error
		db, err = gorm.Open("mysql", dbValue)
		if err != nil {
			log.SQL.Error("failed to connect database, err = " + err.Error())
			return err
		}
		db.LogMode(true)
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		db.DB().SetMaxIdleConns(10)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		db.DB().SetMaxOpenConns(100)
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		db.DB().SetConnMaxLifetime(time.Hour)
		go dbKeepAlive(db)
	}
	return nil
}

func reConnect() error {
	return Connect(dbURLSelf, dbUserSelf, dbPassSelf, dbNameSelf)
}

// Exec 执行自定义 SQL
func Exec(f func(db *gorm.DB)) error {
	if nil == db {
		if err := reConnect(); nil == err {
			f(db)
		} else {
			return err
		}
	}
	return nil
}

// Custom 新建并链接到指定数据库，同时执行 SQL
func Custom(dbURL string, dbUser string, dbPass string, dbName string, f func(db *gorm.DB)) error {
	if err := Connect(dbURL, dbUser, dbPass, dbName); nil == err {
		f(db)
	} else {
		return err
	}
	return nil
}

// Fromat SQL 格式化
func Fromat(elem ...string) string {
	return strings.Join(elem, " ")
}

func dbKeepAlive(db *gorm.DB) {
	c := cron.New()
	_ = c.AddFunc("*/10 * * * * ?", func() {
		err := db.DB().Ping()
		if nil != err {
			_ = Exec(func(db *gorm.DB) {})
		}
	}) //每10秒执行一次
	c.Start()
	select {} //阻塞主线程不退出
}
