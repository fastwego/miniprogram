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

// Package plugin 插件管理
package plugin

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiApplyPlugin           = "/wxa/plugin"
	apiGetPluginDevApplyList = "/wxa/devplugin"
)

/*
向插件开发者发起使用插件的申请/查询已添加的插件/删除已添加的插件



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.applyPlugin.html

POST https://api.weixin.qq.com/wxa/plugin?access_token=TOKEN
*/
func ApplyPlugin(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiApplyPlugin, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取当前所有插件使用方/修改插件使用申请的状态



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginDevApplyList.html

POST https://api.weixin.qq.com/wxa/devplugin?access_token=TOKEN
*/
func GetPluginDevApplyList(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetPluginDevApplyList, bytes.NewReader(payload), "application/json;charset=utf-8")
}
