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

// Package operation 运维中心
package operation

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiGetFeedback       = "/wxaapi/feedback/list"
	apiGetJsErrSearch    = "/wxaapi/log/jserr_search"
	apiGetPerformance    = "/wxaapi/log/get_performance"
	apiGetSceneList      = "/wxaapi/log/get_scene"
	apiGetVersionList    = "/wxaapi/log/get_client_version"
	apiRealtimelogSearch = "/wxaapi/userlog/userlog_search"
)

/*
获取用户反馈列表



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getFeedback.html

GET https://api.weixin.qq.com/wxaapi/feedback/list?access_token=ACCESS_TOKEN
*/
func GetFeedback(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetFeedback)
}

/*
错误查询



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getJsErrSearch.html

POST https://api.weixin.qq.com/wxaapi/log/jserr_search?access_token=ACCESS_TOKEN
*/
func GetJsErrSearch(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetJsErrSearch, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
性能监控



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getPerformance.html

POST https://api.weixin.qq.com/wxaapi/log/get_performance?access_token=ACCESS_TOKEN
*/
func GetPerformance(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetPerformance, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取访问来源



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getSceneList.html

GET https://api.weixin.qq.com/wxaapi/log/get_scene?access_token=ACCESS_TOKEN
*/
func GetSceneList(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetSceneList)
}

/*
获取客户端版本



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getVersionList.html

GET https://api.weixin.qq.com/wxaapi/log/get_client_version?access_token=ACCESS_TOKEN
*/
func GetVersionList(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetVersionList)
}

/*
实时日志查询



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.realtimelogSearch.html

GET https://api.weixin.qq.com/wxaapi/userlog/userlog_search?access_token=ACCESS_TOKEN
*/
func RealtimelogSearch(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiRealtimelogSearch)
}
