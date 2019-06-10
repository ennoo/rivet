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

package sql

import (
	"github.com/ennoo/rivet/utils/env"
	"github.com/ennoo/rivet/utils/log"
	"testing"
)

type User struct {
	Host string `gorm:"column:Host"`
	User string `gorm:"column:User"`
}

func TestSQL(t *testing.T) {
	dbURL := env.GetEnvDefault(env.DBUrl, "127.0.0.1:3306")
	dbPass := env.GetEnvDefault(env.DBPass, "")
	dbName := env.GetEnvDefault(env.DBName, "mysql")
	db := GetSQLInstance()
	_ = db.Connect(dbURL, "root", dbPass, dbName, false)

	_ = db.Connect(dbURL, "root", dbPass, dbName, false)
	log.Common.Info("dbURL = " + db.DBUrl)
	var user User
	db.ExecSQL(&user, "select * from user where User=? limit 1", "root")
	log.Common.Info("user Host = " + user.Host + " User = " + user.User)

	db.DB = nil
	_ = db.reConnect()
	db.ExecSQL(&user, "select * from user where User=? limit 1", "mysql.sys")
	db.DB = nil
	_ = db.Exec(func(sql *SQL) {
		sql.DB.Raw(Format(
			"select * from", "user", "where User=? limit 1"), "mysql.session").Scan(&user)
	})
	_ = db.Exec(func(sql *SQL) {
		sql.DB.Raw(Format(
			"select * from", "user", "where User=? limit 1"), "root").Scan(&user)
	})
	log.Common.Info("user Host = " + user.Host + " User = " + user.User)
}
