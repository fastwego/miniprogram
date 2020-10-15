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

// Package qrcode 普通链接二维码与小程序码
package qrcode

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiQRCodeJumpGet      = "/cgi-bin/wxopen/qrcodejumpget"
	apiQRCodeJumpDownload = "/cgi-bin/wxopen/qrcodejumpdownload"
	apiQRCodeJumpAdd      = "/cgi-bin/wxopen/qrcodejumpadd"
	apiQRCodeJumpPublish  = "/cgi-bin/wxopen/qrcodejumppublish"
	apiQRCodeJumpDelete   = "/cgi-bin/wxopen/qrcodejumpdelete"
	apiShorturl           = "/cgi-bin/shorturl"
	apiGetWxaCodeUnLimit  = "/wxa/getwxacodeunlimit"
	apiGetWxaCode         = "/wxa/getwxacode"
	apiCreateWxaQRCode    = "/cgi-bin/wxaapp/createwxaqrcode"
)

/*
获取已设置的二维码规则

通过本接口可以获取已设置的普通链接二维码规则。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/qrcodejumpget.html

POST https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpget?access_token=ACCESS_TOKEN
*/
func QRCodeJumpGet(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQRCodeJumpGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取校验文件名称及内容

通过本接口下载随机校验文件，并将文件上传至服务器指定位置的目录下，方可通过所属权校验。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/qrcodejumpdownload.html

POST https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpdownload?access_token=ACCESS_TOKEN
*/
func QRCodeJumpDownload(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQRCodeJumpDownload, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
增加或修改二维码规则

通过本接口可以增加或修改普通链接二维码规则.使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/qrcodejumpadd.html

POST https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpadd?access_token=ACCESS_TOKEN
*/
func QRCodeJumpAdd(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQRCodeJumpAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
发布已设置的二维码规则

需要先添加二维码规则，然后调用本接口将二维码规则发布生效，发布后现网用户扫码命中改规则的普通链接二维码时将调整到正式版小程序指定的页面。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/qrcodejumppublish.html

POST https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumppublish?access_token=ACCESS_TOKEN
*/
func QRCodeJumpPublish(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQRCodeJumpPublish, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除已设置的二维码规则

该接口用于删除已设置的二维码规则。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/qrcodejumpdelete.html

POST https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpdelete?access_token=TOKEN
*/
func QRCodeJumpDelete(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQRCodeJumpDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
将二维码长链接转成短链接

调用本 API 可以将一条长链接转成短链接。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/shorturl.html

POST https://api.weixin.qq.com/cgi-bin/shorturl?access_token=ACCESS_TOKEN
*/
func Shorturl(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiShorturl, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取unlimit小程序码

调用本 API 可以获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。 使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/getwxacodeunlimit.html

POST https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN
*/
func GetWxaCodeUnLimit(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetWxaCodeUnLimit, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取小程序码

调用本 API 可以获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/getwxacode.html

POST https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN
*/
func GetWxaCode(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetWxaCode, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取小程序二维码

调用本 API 可以获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/createwxaqrcode.html

POST https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN
*/
func CreateWxaQRCode(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreateWxaQRCode, bytes.NewReader(payload), "application/json;charset=utf-8")
}
