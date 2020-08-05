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

// Package service_market 服务市场
package service_market

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiInvokeService = "/wxa/servicemarket"
)

/*
调用服务平台提供的服务



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/service-market/serviceMarket.invokeService.html

POST https://api.weixin.qq.com/wxa/servicemarket?access_token=ACCESS_TOKEN
*/
func InvokeService(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiInvokeService, bytes.NewReader(payload), "application/json;charset=utf-8")
}
