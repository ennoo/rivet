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

// Package jwt jwt操作工具
package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ennoo/rivet/utils/log"
	"go.uber.org/zap"
)

const (
	// SigningMethodHS256 HS256
	SigningMethodHS256 = iota
	// SigningMethodHS384 HS384
	SigningMethodHS384
	// SigningMethodHS512 HS512
	SigningMethodHS512
)

// Build 创建一个 jwt token
// "sub": "1",  该JWT所面向的用户
// "iss": "http://localhost:8000/user/sign_up", 该JWT的签发者
// "iat": 1451888119, 在什么时候签发的token
// "exp": 1454516119, token什么时候过期
// "nbf": 1451888119, token在此时间之前不能被接收处理
// "jti": "37c107e4609ddbcc9c096ea5ee76c667" token提供唯一标识
func Build(method int, key interface{}, sub, iss, jti string, iat, nbf, exp int64) (string, error) {
	var jwtMethod jwt.SigningMethod
	switch method {
	case SigningMethodHS256:
		jwtMethod = jwt.SigningMethodHS256
	case SigningMethodHS384:
		jwtMethod = jwt.SigningMethodHS384
	case SigningMethodHS512:
		jwtMethod = jwt.SigningMethodHS512
	default:
		jwtMethod = jwt.SigningMethodHS256
	}
	return token(jwtMethod, key, sub, iss, jti, iat, nbf, exp)
}

func token(jwtMethod jwt.SigningMethod, key interface{}, sub, iss, jti string, iat, nbf, exp int64) (tokenString string, err error) {
	token := &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": jwtMethod.Alg(),
		},
		Claims: jwt.StandardClaims{
			Subject:   sub,
			Issuer:    iss,
			Id:        jti,
			IssuedAt:  iat,
			NotBefore: nbf,
			ExpiresAt: exp,
		},
		Method: jwtMethod,
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString(key)

	log.Common.Info("result", zap.String("token", tokenString), zap.Error(err))
	return
}

// Check 验证传入 token 是否合法
func Check(key interface{}, token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Common.Warn("parase with claims failed.", zap.Error(err))
		return false
	}
	return true
}
