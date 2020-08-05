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

// Package mini 物流助手/小程序使用
package mini

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiBatchGetOrder   = "/cgi-bin/express/business/order/batchget"
	apiAddOrder        = "/cgi-bin/express/business/order/add"
	apiBindAccount     = "/cgi-bin/express/business/account/bind"
	apiCancelOrder     = "/cgi-bin/express/business/order/cancel"
	apiGetAllAccount   = "/cgi-bin/express/business/account/getall"
	apiGetAllDelivery  = "/cgi-bin/express/business/delivery/getall"
	apiGetOrder        = "/cgi-bin/express/business/order/get"
	apiGetPath         = "/cgi-bin/express/business/path/get"
	apiGetPrinter      = "/cgi-bin/express/business/printer/getall"
	apiGetQuota        = "/cgi-bin/express/business/quota/get"
	apiTestUpdateOrder = "/cgi-bin/express/business/test_update_order"
	apiUpdatePrinter   = "/cgi-bin/express/business/printer/update"
)

/*
批量获取运单数据



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.batchGetOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/business/order/batchget?access_token=ACCESS_TOKEN
*/
func BatchGetOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchGetOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
生成运单



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.addOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/business/order/add?access_token=ACCESS_TOKEN
*/
func AddOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
绑定、解绑物流账号



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.bindAccount.html

POST https://api.weixin.qq.com/cgi-bin/express/business/account/bind?access_token=ACCESS_TOKEN
*/
func BindAccount(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBindAccount, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
取消运单



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.cancelOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/business/order/cancel?access_token=ACCESS_TOKEN
*/
func CancelOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCancelOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取所有绑定的物流账号



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getAllAccount.html

GET https://api.weixin.qq.com/cgi-bin/express/business/account/getall?access_token=ACCESS_TOKEN
*/
func GetAllAccount(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetAllAccount)
}

/*
获取支持的快递公司列表



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getAllDelivery.html

GET https://api.weixin.qq.com/cgi-bin/express/business/delivery/getall?access_token=ACCESS_TOKEN
*/
func GetAllDelivery(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetAllDelivery)
}

/*
获取运单数据



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/business/order/get?access_token=ACCESS_TOKEN
*/
func GetOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询运单轨迹



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getPath.html

POST https://api.weixin.qq.com/cgi-bin/express/business/path/get?access_token=ACCESS_TOKEN
*/
func GetPath(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetPath, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取打印员。若需要使用微信打单 PC 软件，才需要调用。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getPrinter.html

GET https://api.weixin.qq.com/cgi-bin/express/business/printer/getall?access_token=ACCESS_TOKEN
*/
func GetPrinter(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetPrinter)
}

/*
获取电子面单余额。仅在使用加盟类快递公司时，才可以调用。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getQuota.html

POST https://api.weixin.qq.com/cgi-bin/express/business/quota/get?access_token=ACCESS_TOKEN
*/
func GetQuota(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetQuota, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
模拟快递公司更新订单状态, 该接口只能用户测试



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.testUpdateOrder.html

POST https://api.weixin.qq.com/cgi-bin/express/business/test_update_order?access_token=ACCESS_TOKEN
*/
func TestUpdateOrder(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTestUpdateOrder, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
配置面单打印员，可以设置多个，若需要使用微信打单 PC 软件，才需要调用。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.updatePrinter.html

POST https://api.weixin.qq.com/cgi-bin/express/business/printer/update?access_token=ACCESS_TOKEN
*/
func UpdatePrinter(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdatePrinter, bytes.NewReader(payload), "application/json;charset=utf-8")
}
