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

// Package subscribe 订阅消息
package subscribe

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiAddTemplate                = "/wxaapi/newtmpl/addtemplate"
	apiDeleteTemplate             = "/wxaapi/newtmpl/deltemplate"
	apiGetCategory                = "/wxaapi/newtmpl/getcategory"
	apiGetPubTemplateKeyWordsById = "/wxaapi/newtmpl/getpubtemplatekeywords"
	apiGetPubTemplateTitleList    = "/wxaapi/newtmpl/getpubtemplatetitles"
	apiGetTemplateList            = "/wxaapi/newtmpl/gettemplate"
	apiSend                       = "/cgi-bin/message/subscribe/send"
)

/*
组合模板并添加至帐号下的个人模板库



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.addTemplate.html

POST https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate?access_token=ACCESS_TOKEN
*/
func AddTemplate(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddTemplate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除帐号下的个人模板



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.deleteTemplate.html

POST https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate?access_token=ACCESS_TOKEN
*/
func DeleteTemplate(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDeleteTemplate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取小程序账号的类目



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getCategory.html

GET https://api.weixin.qq.com/wxaapi/newtmpl/getcategory?access_token=ACCESS_TOKEN
*/
func GetCategory(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetCategory)
}

/*
获取模板标题下的关键词列表



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateKeyWordsById.html

GET https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords?access_token=ACCESS_TOKEN
*/
func GetPubTemplateKeyWordsById(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetPubTemplateKeyWordsById)
}

/*
获取帐号所属类目下的公共模板标题



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateTitleList.html

GET https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles?access_token=ACCESS_TOKEN
*/
func GetPubTemplateTitleList(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetPubTemplateTitleList)
}

/*
获取当前帐号下的个人模板列表



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html

GET https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate?access_token=ACCESS_TOKEN
*/
func GetTemplateList(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetTemplateList)
}

/*
发送订阅消息



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html

POST https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=ACCESS_TOKEN
*/
func Send(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSend, bytes.NewReader(payload), "application/json;charset=utf-8")
}
