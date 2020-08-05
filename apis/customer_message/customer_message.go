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

// Package customer_message 客服消息
package customer_message

import (
	"bytes"
	"net/url"

	"github.com/fastwego/miniprogram"
)

const (
	apiGetTempMedia    = "/cgi-bin/media/get"
	apiSend            = "/cgi-bin/message/custom/send"
	apiSetTyping       = "/cgi-bin/message/custom/typing"
	apiUploadTempMedia = "/cgi-bin/media/upload"
)

/*
获取客服消息内的临时素材。即下载临时的多媒体文件。目前小程序仅支持下载图片文件。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.getTempMedia.html

GET https://api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
*/
func GetTempMedia(ctx *miniprogram.Miniprogram, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetTempMedia + "?" + params.Encode())
}

/*
发送客服消息给用户。详细规则见 发送客服消息



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html

POST https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
*/
func Send(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSend, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
下发客服当前输入状态给用户。详见 客服消息输入状态



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html

POST https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN
*/
func SetTyping(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetTyping, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
把媒体文件上传到微信服务器。目前仅支持图片。用于发送客服消息或被动回复用户消息。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html

POST https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
*/
func UploadTempMedia(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUploadTempMedia+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}
