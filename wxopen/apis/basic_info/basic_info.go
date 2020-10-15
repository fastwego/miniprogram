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

// Package basic_info 基础信息设置
package basic_info

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiGetAccountBasicInfo   = "/cgi-bin/account/getaccountbasicinfo"
	apiModifyDomain          = "/wxa/modify_domain"
	apiSetWebviewDomain      = "/wxa/setwebviewdomain"
	apiSetNickname           = "/wxa/setnickname"
	apiCheckWXVerifyNickname = "/cgi-bin/wxverify/checkwxverifynickname"
	apiQueryNickname         = "/wxa/api_wxa_querynickname"
	apiModifyHeadImage       = "/cgi-bin/account/modifyheadimage"
	apiModifySignature       = "/cgi-bin/account/modifysignature"
	apiGetWxaSearchStatus    = "/wxa/getwxasearchstatus"
	apiChangeWxaSearchStatus = "/wxa/changewxasearchstatus"
)

/*
获取基本信息

调用本 API 可以获取小程序的基本信息。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Mini_Program_Information_Settings.html

GET https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?access_token=ACCESS_TOKEN
*/
func GetAccountBasicInfo(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetAccountBasicInfo)
}

/*
设置服务器域名

授权给第三方的小程序，其服务器域名只可以为在第三方平台账号中配置的小程序服务器域名，当小程序通过第三方平台发布代码上线后，小程序原先自己配置的服务器域名将被删除，只保留第三方平台的域名，所以第三方平台在代替小程序发布代码之前，需要调用接口为小程序添加第三方平台自身的域名。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Server_Address_Configuration.html

POST https://api.weixin.qq.com/wxa/modify_domain?access_token=ACCESS_TOKEN
*/
func ModifyDomain(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiModifyDomain, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置业务域名

授权给第三方的小程序，其业务域名只可以为在第三方平台账号中配置的小程序业务域名，当小程序通过第三方发布代码上线后，小程序原先自己配置的业务域名将被删除，只保留第三方平台的域名，所以第三方平台在代替小程序发布代码之前，需要调用接口为小程序添加业务域名。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/setwebviewdomain.html

POST https://api.weixin.qq.com/wxa/setwebviewdomain?access_token=ACCESS_TOKEN
*/
func SetWebviewDomain(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetWebviewDomain, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置名称

调用本接口可以设置小程序名称，当名称没有命中关键词，则直接设置成功；当名称命中关键词，需提交证明材料，并需要审核。审核结果会向消息与事件接收 URL 进行事件推送。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/setnickname.html

POST https://api.weixin.qq.com/wxa/setnickname?access_token=ACCESS_TOKEN
*/
func SetNickname(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetNickname, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
微信认证名称检测

调用本 API 可以检测微信认证的名称是否符合规则。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。POST https://api.weixin.qq.com/cgi-bin/wxverify/checkwxverifynickname?access_token=ACCESS_TOKEN

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/wxverify_checknickname.html

POST https://api.weixin.qq.com/cgi-bin/wxverify/checkwxverifynickname?access_token=ACCESS_TOKEN
*/
func CheckWXVerifyNickname(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCheckWXVerifyNickname, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询改名审核状态

调用设置名称接口，如果需要审核，会返回审核单 id（audit_id）,使用本接口可以查询改名审核状态。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/api_wxa_querynickname.html

POST https://api.weixin.qq.com/wxa/api_wxa_querynickname?access_token=ACCESS_TOKEN
*/
func QueryNickname(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQueryNickname, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改头像

调用本接口可以修改小程序的头像。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/modifyheadimage.html

POST https://api.weixin.qq.com/cgi-bin/account/modifyheadimage?access_token=ACCESS_TOKEN
*/
func ModifyHeadImage(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiModifyHeadImage, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改功能介绍

调用本接口可以修改功能介绍。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/modifysignature.html

POST https://api.weixin.qq.com/cgi-bin/account/modifysignature?access_token=ACCESS_TOKEN
*/
func ModifySignature(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiModifySignature, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询隐私设置

通过本接口可以查询小程序当前的隐私设置，即是否可被搜索。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/getwxasearchstatus.html

GET https://api.weixin.qq.com/wxa/getwxasearchstatus?access_token=ACCESS_TOKEN
*/
func GetWxaSearchStatus(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetWxaSearchStatus)
}

/*
修改隐私设置

通过本接口修改小程序隐私设置，即修改是否可被搜索，可以先查询小程序当前的隐私设置再决定是否修改。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/changewxasearchstatus.html

POST https://api.weixin.qq.com/wxa/changewxasearchstatus?access_token=ACCESS_TOKEN
*/
func ChangeWxaSearchStatus(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiChangeWxaSearchStatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}
