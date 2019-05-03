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

package dolphin

import (
	"encoding/json"
	"github.com/ennoo/rivet/common/util/log"
	"github.com/ennoo/rivet/dolphin/http/request"
	"github.com/ennoo/rivet/dolphin/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(context *gin.Context, remote string, uri string) {
	res := response.Result{}
	req := context.Request
	restTransHandler := request.RestTransHandler{
		RestHandler: request.RestHandler{
			RemoteServer: remote,
			Uri:          uri,
			Body:         req.Body,
			Header:       nil,
			Cookies:      nil}}
	body, err := restTransHandler.Get()
	if err != nil {
		res.Fail(err.Error())
	} else {
		log.Debug("body = ", string(body))
		err := json.Unmarshal(body, &res)
		if nil != err {
			res.Fail(err.Error())
		}
	}
	context.JSON(http.StatusOK, res)
}
