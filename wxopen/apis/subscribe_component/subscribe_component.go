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

// Package subscribe_component 扫码关注组件
package subscribe_component

import (
	"bytes"
	"net/url"

	"github.com/fastwego/miniprogram"
)

const (
	apiGetShowWxaItem      = "/wxa/getshowwxaitem"
	apiGetWxaMplinkForShow = "/wxa/getwxamplinkforshow"
	apiUpdateShowWxaItem   = "/wxa/updateshowwxaitem"
)

/*
获取展示的公众号信息

使用本接口可以获取扫码关注组件所展示的公众号信息

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_component/getshowwxaitem.html

GET https://api.weixin.qq.com/wxa/getshowwxaitem?access_token=ACCESS_TOKEN
*/
func GetShowWxaItem(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetShowWxaItem)
}

/*
获取可以用来设置的公众号列表

通过本接口可以获取扫码关注组件允许展示的公众号列表

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_component/getwxamplinkforshow.html

GET https://api.weixin.qq.com/wxa/getwxamplinkforshow?page=0&num=20&access_token=ACCESS_TOKEN
*/
func GetWxaMplinkForShow(ctx *miniprogram.Miniprogram, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetWxaMplinkForShow + "?" + params.Encode())
}

/*
设置展示的公众号信息

使用本接口可以设置扫码关注组件所展示的公众号信息

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_component/updateshowwxaitem.html

POST https://api.weixin.qq.com/wxa/updateshowwxaitem?access_token=ACCESS_TOKEN
*/
func UpdateShowWxaItem(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateShowWxaItem, bytes.NewReader(payload), "application/json;charset=utf-8")
}
