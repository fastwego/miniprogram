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

// Package data 开放数据
package data

import (
	"bytes"
	"net/url"

	"github.com/fastwego/miniprogram"
)

const (
	apiRemoveUserStorage      = "/wxa/remove_user_storage"
	apiSetUserInteractiveData = "/wxa/setuserinteractivedata"
	apiSetUserStorage         = "/wxa/set_user_storage"
)

/*
删除已经上报到微信的key-value数据



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data/storage.removeUserStorage.html

POST https://api.weixin.qq.com/wxa/remove_user_storage?access_token=ACCESS_TOKEN&signature=SIGNATURE&openid=OPENID&sig_method=SIG_METHOD
*/
func RemoveUserStorage(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiRemoveUserStorage+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
写用户关系链互动数据存储



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data/storage.setUserInteractiveData.html

POST https://api.weixin.qq.com/wxa/setuserinteractivedata?access_token=ACCESS_TOKEN&signature=SIGNATURE&openid=OPENID&sig_method=SIG_METHOD
*/
func SetUserInteractiveData(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetUserInteractiveData+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
上报用户数据后台接口。小游戏可以通过本接口上报key-value数据到用户的CloudStorage。



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data/storage.setUserStorage.html

POST https://api.weixin.qq.com/wxa/set_user_storage?access_token=ACCESS_TOKEN&signature=SIGNATURE&openid=OPENID&sig_method=SIG_METHOD
*/
func SetUserStorage(ctx *miniprogram.Miniprogram, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetUserStorage+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}
