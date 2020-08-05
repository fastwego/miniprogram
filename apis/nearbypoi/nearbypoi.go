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

// Package nearbypoi 附近的小程序
package nearbypoi

import (
	"bytes"
	"net/url"

	"github.com/fastwego/miniprogram"
)

const (
	apiAdd           = "/wxa/addnearbypoi"
	apiDelete        = "/wxa/delnearbypoi"
	apiGetList       = "/wxa/getnearbypoilist"
	apiSetShowStatus = "/wxa/setnearbypoishowstatus"
)

/*
添加地点



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.add.html

POST https://api.weixin.qq.com/wxa/addnearbypoi?access_token=ACCESS_TOKEN
*/
func Add(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除地点



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.delete.html

POST https://api.weixin.qq.com/wxa/delnearbypoi?access_token=ACCESS_TOKEN
*/
func Delete(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查看地点列表



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.getList.html

GET https://api.weixin.qq.com/wxa/getnearbypoilist?page=1&page_rows=20&access_token=ACCESS_TOKEN
*/
func GetList(ctx *miniprogram.Miniprogram, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetList + "?" + params.Encode())
}

/*
展示/取消展示附近小程序



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.setShowStatus.html

POST https://api.weixin.qq.com/wxa/setnearbypoishowstatus?access_token=ACCESS_TOKEN
*/
func SetShowStatus(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetShowStatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}
