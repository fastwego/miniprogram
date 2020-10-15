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

// Package code_management 代码管理
package code_management

import (
	"bytes"
	"net/url"

	"github.com/fastwego/miniprogram"
)

const (
	apiCommit                 = "/wxa/commit"
	apiGetPage                = "/wxa/get_page"
	apiGetQrcode              = "/wxa/get_qrcode"
	apiSubmitAudit            = "/wxa/submit_audit"
	apiGetAuditstatus         = "/wxa/get_auditstatus"
	apiGetLatestAuditstatus   = "/wxa/get_latest_auditstatus"
	apiUndoCodeAudit          = "/wxa/undocodeaudit"
	apiRelease                = "/wxa/release"
	apiRevertCodeRelease      = "/wxa/revertcoderelease"
	apiGrayRelease            = "/wxa/grayrelease"
	apiGetGrayReleasePlan     = "/wxa/getgrayreleaseplan"
	apiRevertGrayRelease      = "/wxa/revertgrayrelease"
	apiChangeVisitstatus      = "/wxa/change_visitstatus"
	apiGetWeAppSupportVersion = "/cgi-bin/wxopen/getweappsupportversion"
	apiSetWeAppSupportVersion = "/cgi-bin/wxopen/setweappsupportversion"
	apiQueryQuota             = "/wxa/queryquota"
	apiSpeedupAudit           = "/wxa/speedupaudit"
)

/*
上传小程序代码

第三方平台需要先将草稿添加到代码模板库，或者从代码模板库中选取某个代码模板，得到对应的模板 id（template_id）； 然后调用本接口可以为已授权的小程序上传代码。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/commit.html

POST https://api.weixin.qq.com/wxa/commit?access_token=ACCESS_TOKEN
*/
func Commit(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCommit, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取已上传的代码的页面列表

通过本接口可以获取由第三方平台上传小程序代码的页面列表；用于提交审核的审核项 的 address 参数使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_page.html

GET https://api.weixin.qq.com/wxa/get_page?access_token=ACCESS_TOKEN
*/
func GetPage(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetPage)
}

/*
获取体验版二维码

调用本接口可以获取小程序的体验版二维码。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_qrcode.html

GET https://api.weixin.qq.com/wxa/get_qrcode?access_token=ACCESS_TOKEN&path=page%2Findex%3Faction%3D1
*/
func GetQrcode(ctx *miniprogram.Miniprogram, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetQrcode + "?" + params.Encode())
}

/*
提交审核

在调用上传代码接口为小程序上传代码后，可以调用本接口，将上传的代码提交审核。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/submit_audit.html

POST https://api.weixin.qq.com/wxa/submit_audit?access_token=ACCESS_TOKEN
*/
func SubmitAudit(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSubmitAudit, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询指定发布审核单的审核状态

提交审核后，调用本接口可以查询指定发布审核单的审核状态。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_auditstatus.html

POST https://api.weixin.qq.com/wxa/get_auditstatus?access_token=ACCESS_TOKEN
*/
func GetAuditstatus(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetAuditstatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询最新一次提交的审核状态

调用本接口可以查询最新一次提审单的审核状态。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_latest_auditstatus.html

GET https://api.weixin.qq.com/wxa/get_latest_auditstatus?access_token=ACCESS_TOKEN
*/
func GetLatestAuditstatus(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetLatestAuditstatus)
}

/*
小程序审核撤回

调用本接口可以撤回当前的代码审核单。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/undocodeaudit.html

GET https://api.weixin.qq.com/wxa/undocodeaudit?access_token=ACCESS_TOKEN
*/
func UndoCodeAudit(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiUndoCodeAudit)
}

/*
发布已通过审核的小程序

调用本接口可以发布最后一个审核通过的小程序代码版本。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/release.html

POST https://api.weixin.qq.com/wxa/release?access_token=ACCESS_TOKEN
*/
func Release(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiRelease, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
版本回退

调用本接口可以将小程序的线上版本进行回退。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/revertcoderelease.html

GET https://api.weixin.qq.com/wxa/revertcoderelease?access_token=ACCESS_TOKEN
*/
func RevertCodeRelease(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiRevertCodeRelease)
}

/*
分阶段发布

发布小程序接口 是全量发布，会影响到现网的所有用户。而本接口是创建一个灰度发布的计划，可以控制发布的节奏，避免一上线就影响到所有的用户。可以多次调用本次接口，将灰度的比例（gray_percentage）逐渐增大使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/grayrelease.html

POST https://api.weixin.qq.com/wxa/grayrelease?access_token=ACCESS_TOKEN
*/
func GrayRelease(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGrayRelease, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询当前分阶段发布详情

该接口用于查询当前分阶段发布详情。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/getgrayreleaseplan.html

GET https://api.weixin.qq.com/wxa/getgrayreleaseplan?access_token=ACCESS_TOKEN
*/
func GetGrayReleasePlan(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetGrayReleasePlan)
}

/*
取消分阶段发布

在小程序分阶段发布期间，可以随时调用本接口取消分阶段发布。取消分阶段发布后，受影响的微信用户（即被灰度升级的微信用户）的小程序版本将回退到分阶段发布前的版本

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/revertgrayrelease.html

GET https://api.weixin.qq.com/wxa/revertgrayrelease?access_token=ACCESS_TOKEN
*/
func RevertGrayRelease(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiRevertGrayRelease)
}

/*
修改小程序线上代码的可见状态



See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/change_visitstatus.html

POST https://api.weixin.qq.com/wxa/change_visitstatus?access_token=ACCESS_TOKEN
*/
func ChangeVisitstatus(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiChangeVisitstatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询当前设置的最低基础库版本及各版本用户占比

调用本接口可以查询小程序当前设置的最低基础库版本，以及小程序在各个基础库版本的用户占比。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/getweappsupportversion.html

POST https://api.weixin.qq.com/cgi-bin/wxopen/getweappsupportversion?access_token=ACCESS_TOKEN
*/
func GetWeAppSupportVersion(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetWeAppSupportVersion, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置最低基础库版本

调用本接口可以设置小程序的最低基础库支持版本，可以先查询当前小程序在各个基础库的用户占比来辅助进行决策使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/setweappsupportversion.html

POST https://api.weixin.qq.com/cgi-bin/wxopen/setweappsupportversion?access_token=ACCESS_TOKEN
*/
func SetWeAppSupportVersion(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetWeAppSupportVersion, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询服务商的当月提审限额（quota）和加急次数

服务商可以调用该接口，查询当月平台分配的提审限额和剩余可提审次数，以及当月分配的审核加急次数和剩余加急次数。（所有旗下小程序共用该额度）使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/query_quota.html

GET https://api.weixin.qq.com/wxa/queryquota?access_token=ACCESS_TOKEN
*/
func QueryQuota(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiQueryQuota)
}

/*
加急审核申请

有加急次数的第三方可以通过该接口，对已经提审的小程序进行加急操作，加急后的小程序预计2-12小时内审完。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/speedup_audit.html

POST https://api.weixin.qq.com/wxa/speedupaudit?access_token=ACCESS_TOKEN
*/
func SpeedupAudit(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpeedupAudit, bytes.NewReader(payload), "application/json;charset=utf-8")
}
