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

// Package wxacode 小程序码
package wxacode

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiCreateQRCode = "/cgi-bin/wxaapp/createwxaqrcode"
	apiGet          = "/wxa/getwxacode"
	apiGetUnlimited = "/wxa/getwxacodeunlimit"
)

/*
获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html

POST https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN
*/
func CreateQRCode(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreateQRCode, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/qr-code/wxacode.get.html

POST https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN
*/
func Get(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。 更多用法详见 获取二维码。



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html

POST https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN
*/
func GetUnlimited(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUnlimited, bytes.NewReader(payload), "application/json;charset=utf-8")
}
