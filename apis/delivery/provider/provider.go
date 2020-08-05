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

// Package provider 即时配送/运力方
package provider

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiUpdateOrder = "/cgi-bin/express/local/delivery/update_order"
)

/*
配送公司更新配送单状态



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.updateOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/local/delivery/update_order?access_token=ACCESS_TOKEN
*/
func UpdateOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}
