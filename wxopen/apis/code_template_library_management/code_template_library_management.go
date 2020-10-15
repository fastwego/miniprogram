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

// Package code_template_library_management 代码模板库设置
package code_template_library_management

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiGetTemplateDraftList = "/wxa/gettemplatedraftlist"
	apiAddToTemplate        = "/wxa/addtotemplate"
	apiGetTemplateList      = "/wxa/gettemplatelist"
	apiDeleteTemplate       = "/wxa/deletetemplate"
)

/*
获取代码草稿列表

通过本接口，可以获取草稿箱中所有的草稿（临时代码模板）；草稿是由第三方平台的开发小程序在使用微信开发者工具上传的，点击查ext详情使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatedraftlist.html

GET https://api.weixin.qq.com/wxa/gettemplatedraftlist?access_token=ACCESS_TOKEN
*/
func GetTemplateDraftList(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetTemplateDraftList)
}

/*
将草稿添加到代码模板库

可以通过获取草稿箱中所有的草稿得到草稿 ID；调用本接口可以将临时草稿选为持久的代码模板。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/addtotemplate.html

POST https://api.weixin.qq.com/wxa/addtotemplate?access_token=TOKEN
*/
func AddToTemplate(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddToTemplate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取代码模板列表

第三方平台运营者可以登录 open.weixin.qq.com 或者通过将草稿箱的草稿选为代码模板接口，将草稿箱中的某个代码版本添加到代码模板库中使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatelist.html

GET https://api.weixin.qq.com/wxa/gettemplatelist?access_token=ACCESS_TOKEN
*/
func GetTemplateList(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetTemplateList)
}

/*
删除指定代码模板

因为代码模板库的模板数量是有上限的，当达到上限或者有某个模板不再需要时，可以调用本接口删除指定的代码模板。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/deletetemplate.html

POST https://api.weixin.qq.com/wxa/deletetemplate?access_token=ACCESS_TOKEN
*/
func DeleteTemplate(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDeleteTemplate, bytes.NewReader(payload), "application/json;charset=utf-8")
}
