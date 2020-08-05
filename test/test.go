// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package test 模拟微信服务器 测试
package test

import (
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/fastwego/miniprogram"
)

var MockMiniprogram *miniprogram.Miniprogram
var MockSvr *httptest.Server
var MockSvrHandler *http.ServeMux
var onceSetup sync.Once

// 初始化测试环境
func Setup() {
	onceSetup.Do(func() {

		MockMiniprogram = miniprogram.New(miniprogram.MiniprogramConfig{
			Appid:  "APPID",
			Secret: "SECRET",
		})

		// Mock Server
		MockSvrHandler = http.NewServeMux()
		MockSvr = httptest.NewServer(MockSvrHandler)
		miniprogram.WXServerUrl = MockSvr.URL // 拦截发往微信服务器的请求

		// Mock access token
		MockSvrHandler.HandleFunc("/cgi-bin/token", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"access_token":"ACCESS_TOKEN","expires_in":7200}`))
		})
	})
}
