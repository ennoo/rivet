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

package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ennoo/rivet/utils/log"
	"go.uber.org/zap"
)

var key []byte = []byte("Hello World！This is secret!")

func Build(sub, iss, jti string, iat, exp, nbf int64) (string, error) {
	// "sub": "1",  该JWT所面向的用户
	// "iss": "http://localhost:8000/user/sign_up", 该JWT的签发者
	// "iat": 1451888119, 在什么时候签发的token
	// "exp": 1454516119, token什么时候过期
	// "nbf": 1451888119, token在此时间之前不能被接收处理
	// "jti": "37c107e4609ddbcc9c096ea5ee76c667" token提供唯一标识
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   sub,
		Issuer:    iss,
		Id:        jti,
		IssuedAt:  iat,
		NotBefore: nbf,
		ExpiresAt: exp,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(key)

	log.Common.Info("result", zap.String("token", tokenString), zap.Error(err))
	return tokenString, err
}

func Check(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte("userMD5"), nil
	})
	if err != nil {
		log.Common.Info("parase with claims failed.", zap.Error(err))
		return false
	}
	return true
}
