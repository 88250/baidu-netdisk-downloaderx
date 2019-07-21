// Gulu - Golang common utilities for everyone.
// Copyright (c) 2019-present, b3log.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gulu

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
)

// Result represents a common-used result struct.
type Result struct {
	Code int         `json:"code"` // return code
	Msg  string      `json:"msg"`  // message
	Data interface{} `json:"data"` // data object
}

// NewResult creates a result with Code=0, Msg="", Data=nil.
func (*GuluRet) NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}

// RetResult writes HTTP response with "Content-Type, application/json".
func (*GuluRet) RetResult(w http.ResponseWriter, r *http.Request, res *Result) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(res)
	if err != nil {
		logger.Error(err)

		return
	}

	w.Write(data)
}

// RetGzResult writes HTTP response with "Content-Type, application/json" and "Content-Encoding, gzip".
func (*GuluRet) RetGzResult(w http.ResponseWriter, r *http.Request, res *Result) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")

	gz := gzip.NewWriter(w)
	err := json.NewEncoder(gz).Encode(res)
	if nil != err {
		logger.Error(err)

		return
	}

	err = gz.Close()
	if nil != err {
		logger.Error(err)

		return
	}
}

// RetJSON writes HTTP response with "Content-Type, application/json".
func (*GuluRet) RetJSON(w http.ResponseWriter, r *http.Request, res map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(res)
	if err != nil {
		logger.Error(err)

		return
	}

	w.Write(data)
}

// RetGzJSON writes HTTP response with "Content-Type, application/json" and "Content-Encoding, gzip".
func (*GuluRet) RetGzJSON(w http.ResponseWriter, r *http.Request, res map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")

	gz := gzip.NewWriter(w)
	err := json.NewEncoder(gz).Encode(res)
	if nil != err {
		logger.Error(err)

		return
	}

	err = gz.Close()
	if nil != err {
		logger.Error(err)

		return
	}
}
