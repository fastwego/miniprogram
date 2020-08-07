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

// Package updatable_message 动态消息
package updatable_message

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiCreateActivityId = "/cgi-bin/message/wxopen/activityid/create"
	apiSetUpdatableMsg  = "/cgi-bin/message/wxopen/updatablemsg/send"
)

/*
创建被分享动态消息的 activity_id。详见动态消息。



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/updatable-message/updatableMessage.createActivityId.html

GET https://api.weixin.qq.com/cgi-bin/message/wxopen/activityid/create?access_token=ACCESS_TOKEN
*/
func CreateActivityId(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiCreateActivityId)
}

/*
修改被分享的动态消息。详见动态消息。



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/updatable-message/updatableMessage.setUpdatableMsg.html

POST https://api.weixin.qq.com/cgi-bin/message/wxopen/updatablemsg/send?access_token=ACCESS_TOKEN
*/
func SetUpdatableMsg(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetUpdatableMsg, bytes.NewReader(payload), "application/json;charset=utf-8")
}
