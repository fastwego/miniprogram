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

// Package frame 帧同步
package frame

import (
	"bytes"
	"net/url"

	"github.com/fastwego/miniprogram"
)

const (
	apiCreateGameRoom      = "/wxa/createwxagameroom"
	apiGetGameFrame        = "/wxa/getwxagameframe"
	apiGetGameIdentityInfo = "/wxa/getwxagameidentityinfo"
	apiGetGameRoomInfo     = "/wxa/getwxagameroominfo"
)

/*
第三方后台创建帧同步游戏房间



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.createGameRoom.html

POST https://api.weixin.qq.com/wxa/createwxagameroom?access_token=ACCESS_TOKEN
*/
func CreateGameRoom(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreateGameRoom, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
分片拉取对局游戏帧



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.getGameFrame.html

GET https://api.weixin.qq.com/wxa/getwxagameframe?access_token=XXX&access_info=YYY&begin_frame_id=ZZZ&end_frame_id=TTT
*/
func GetGameFrame(ctx *miniprogram.Miniprogram, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetGameFrame + "?" + params.Encode())
}

/*
获取对局玩家位次信息



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.getGameIdentityInfo.html

GET https://api.weixin.qq.com/wxa/getwxagameidentityinfo?access_token=XXX&access_info=YYY
*/
func GetGameIdentityInfo(ctx *miniprogram.Miniprogram, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetGameIdentityInfo + "?" + params.Encode())
}

/*
获取指定房间信息



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.getGameRoomInfo.html

GET https://api.weixin.qq.com/wxa/getwxagameroominfo?access_token=XXX&access_info=YYY
*/
func GetGameRoomInfo(ctx *miniprogram.Miniprogram, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetGameRoomInfo + "?" + params.Encode())
}
