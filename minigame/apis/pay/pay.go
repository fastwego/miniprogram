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

// Package pay 虚拟支付
package pay

import (
	"bytes"
	"fmt"

	"github.com/fastwego/miniprogram"
)

const (
	apiCancelPay  = "/cgi-bin/midas%s/cancelpay"
	apiGetBalance = "/cgi-bin/midas%s/getbalance"
	apiPay        = "/cgi-bin/midas%s/pay"
	apiPresent    = "/cgi-bin/midas%s/present"

	apiGetOrder = "/wxa/business/getpayfororder"
)

// 开启沙箱模式
func setSandBoxMode(ctx *miniprogram.Miniprogram, apiUri string) string {
	if ctx.Config.PaySandbox {
		return fmt.Sprintf(apiUri, "/sandbox")
	}
	return fmt.Sprintf(apiUri, "")
}

/*
取消订单。开通了虚拟支付的小游戏，若扣除游戏币的订单号在有效时间内，可以通过本接口取消该笔扣除游戏币的订单

See: https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.cancelPay.html

POST https://api.weixin.qq.com/cgi-bin/midas/cancelpay?access_token=ACCESS_TOKEN
*/
func CancelPay(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(setSandBoxMode(ctx, apiCancelPay), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取游戏币余额。开通了虚拟支付的小游戏，可以通过本接口查看某个用户的游戏币余额

See: https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.getBalance.html

POST https://api.weixin.qq.com/cgi-bin/midas/getbalance?access_token=ACCESS_TOKEN
*/
func GetBalance(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(setSandBoxMode(ctx, apiGetBalance), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
扣除游戏币。开通了虚拟支付的小游戏，可以通过本接口扣除某个用户的游戏币。 由于可能存在接口调用超时或返回系统失败，但是游戏币实际已经扣除的情况，所以当该接口返回系统失败时，可以用相同的bill_no再次调用本接口，直到返回非系统失败为止，不会重复扣款，也可以调用取消支付接口取消本次扣款。

See: https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.pay.html

POST https://api.weixin.qq.com/cgi-bin/midas/pay?access_token=ACCESS_TOKEN
*/
func Pay(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(setSandBoxMode(ctx, apiPay), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
给用户赠送游戏币。开通了虚拟支付的小游戏，可以通过该接口赠送游戏币给某个用户。

See: https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.present.html

POST https://api.weixin.qq.com/cgi-bin/midas/present?access_token=ACCESS_TOKEN
*/
func Present(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(setSandBoxMode(ctx, apiPresent), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询索要单。开发者可以通过查询接口主动查询索要单信息和状态，完成下一步的业务逻辑

See: https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/business.getPayForOrder.html

POST https://api.weixin.qq.com/wxa/business/getpayfororder?access_token=ACCESS_TOKEN
*/
func GetPayForOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}
