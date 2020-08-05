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

// Package mini 即时配送/小程序使用
package mini

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiAbnormalConfirm     = "/cgi-bin/express/local/business/order/confirm_return"
	apiAddOrder            = "/cgi-bin/express/local/business/order/add"
	apiAddTip              = "/cgi-bin/express/local/business/order/addtips"
	apiBindAccount         = "/cgi-bin/express/local/business/shop/add"
	apiCancelOrder         = "/cgi-bin/express/local/business/order/cancel"
	apiGetAllImmeDelivery  = "/cgi-bin/express/local/business/delivery/getall"
	apiGetBindAccount      = "/cgi-bin/express/local/business/shop/get"
	apiGetOrder            = "/cgi-bin/express/local/business/order/get"
	apiMockUpdateOrder     = "/cgi-bin/express/local/business/test_update_order"
	apiOpenDelivery        = "/cgi-bin/express/local/business/open"
	apiPreAddOrder         = "/cgi-bin/express/local/business/order/pre_add"
	apiPreCancelOrder      = "/cgi-bin/express/local/business/order/precancel"
	apiRealMockUpdateOrder = "/cgi-bin/express/local/business/realmock_update_order"
	apiReOrder             = "/cgi-bin/express/local/business/order/readd"
)

/*
异常件退回商家商家确认收货接口



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.abnormalConfirm.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/confirm_return?access_token=ACCESS_TOKEN
*/
func AbnormalConfirm(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAbnormalConfirm, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
下配送单接口



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/add?access_token=ACCESS_TOKEN
*/
func AddOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
可以对待接单状态的订单增加小费。需要注意：订单的小费，以最新一次加小费动作的金额为准，故下一次增加小费额必须大于上一次小费额



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addTip.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/addtips?access_token=ACCESS_TOKEN
*/
func AddTip(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddTip, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
第三方代商户发起绑定配送公司帐号的请求



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.bindAccount.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/shop/add?access_token=COMPNENT_ACCESS_TOKEN
*/
func BindAccount(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBindAccount, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
取消配送单接口



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.cancelOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/cancel?access_token=ACCESS_TOKEN
*/
func CancelOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCancelOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取已支持的配送公司列表接口



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getAllImmeDelivery.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/delivery/getall?access_token=ACCESS_TOKEN
*/
func GetAllImmeDelivery(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetAllImmeDelivery, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
拉取已绑定账号



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getBindAccount.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/shop/get?access_token=ACCESS_TOKEN
*/
func GetBindAccount(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetBindAccount, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
拉取配送单信息



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/get?access_token=ACCESS_TOKEN
*/
func GetOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
模拟配送公司更新配送单状态, 该接口只用于沙盒环境，即订单并没有真实流转到运力方



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.mockUpdateOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/test_update_order?access_token=ACCESS_TOKEN
*/
func MockUpdateOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMockUpdateOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
第三方代商户发起开通即时配送权限



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.openDelivery.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/open?access_token=COMPNENT_ACCESS_TOKEN
*/
func OpenDelivery(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOpenDelivery, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
预下配送单接口



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.preAddOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/pre_add?access_token=ACCESS_TOKEN
*/
func PreAddOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPreAddOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
预取消配送单接口



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.preCancelOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/precancel?access_token=ACCESS_TOKEN
*/
func PreCancelOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPreCancelOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
模拟配送公司更新配送单状态, 该接口用于测试账户下的单，将请求转发到运力测试环境



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.realMockUpdateOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/realmock_update_order?access_token=ACCESS_TOKEN
*/
func RealMockUpdateOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiRealMockUpdateOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
重新下单



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.reOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/readd?access_token=ACCESS_TOKEN
*/
func ReOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiReOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}
