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

// Package ocr OCR
package ocr

import (
	"bytes"
	"net/url"

	"github.com/fastwego/miniprogram"
)

const (
	apiBankcard        = "/cv/ocr/bankcard"
	apiBusinessLicense = "/cv/ocr/bizlicense"
	apiDriverLicense   = "/cv/ocr/drivinglicense"
	apiIdcard          = "/cv/ocr/idcard"
	apiPrintedText     = "/cv/ocr/comm"
	apiVehicleLicense  = "/cv/ocr/driving"
)

/*
本接口提供基于小程序的银行卡 OCR 识别



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.bankcard.html

POST https://api.weixin.qq.com/cv/ocr/bankcard?type=MODE&img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func Bankcard(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBankcard+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
本接口提供基于小程序的营业执照 OCR 识别



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.businessLicense.html

POST https://api.weixin.qq.com/cv/ocr/bizlicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func BusinessLicense(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBusinessLicense+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
本接口提供基于小程序的驾驶证 OCR 识别



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.driverLicense.html

POST https://api.weixin.qq.com/cv/ocr/drivinglicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func DriverLicense(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDriverLicense+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
本接口提供基于小程序的身份证 OCR 识别



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.idcard.html

POST https://api.weixin.qq.com/cv/ocr/idcard?type=MODE&img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func Idcard(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiIdcard+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
本接口提供基于小程序的通用印刷体 OCR 识别



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.printedText.html

POST https://api.weixin.qq.com/cv/ocr/comm?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func PrintedText(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPrintedText+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
本接口提供基于小程序的行驶证 OCR 识别



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.vehicleLicense.html

POST https://api.weixin.qq.com/cv/ocr/driving?type=MODE&img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func VehicleLicense(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiVehicleLicense+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}
