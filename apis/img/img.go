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

// Package img 图像处理
package img

import (
	"bytes"
	"net/url"

	"github.com/fastwego/miniprogram"
)

const (
	apiAiCrop          = "/cv/img/aicrop"
	apiScanQRCode      = "/cv/img/qrcode"
	apiSuperresolution = "/cv/img/superresolution"
)

/*
本接口提供基于小程序的图片智能裁剪能力。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.aiCrop.html

POST https://api.weixin.qq.com/cv/img/aicrop?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func AiCrop(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAiCrop+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
本接口提供基于小程序的条码/二维码识别的API。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.scanQRCode.html

POST https://api.weixin.qq.com/cv/img/qrcode?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func ScanQRCode(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiScanQRCode+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
本接口提供基于小程序的图片高清化能力。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.superresolution.html

POST https://api.weixin.qq.com/cv/img/superresolution?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func Superresolution(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSuperresolution+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}
