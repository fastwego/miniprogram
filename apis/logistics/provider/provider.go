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

// Package provider 物流助手/运力方
package provider

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiGetContact      = "/cgi-bin/express/delivery/contact/get"
	apiPreviewTemplate = "/cgi-bin/express/delivery/template/preview"
	apiUpdateBusiness  = "/cgi-bin/express/delivery/service/business/update"
	apiUpdatePath      = "/cgi-bin/express/delivery/path/update"
)

/*
获取面单联系人信息



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.getContact.html

POST https://api.weixin.qq.com/cgi-bin/express/delivery/contact/get?access_token=ACCESS_TOKEN
*/
func GetContact(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetContact, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
预览面单模板。用于调试面单模板使用。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.previewTemplate.html

POST https://api.weixin.qq.com/cgi-bin/express/delivery/template/preview?access_token=ACCESS_TOKEN
*/
func PreviewTemplate(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPreviewTemplate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新商户审核结果



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.updateBusiness.html

POST https://api.weixin.qq.com/cgi-bin/express/delivery/service/business/update?access_token=ACCESS_TOKEN
*/
func UpdateBusiness(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateBusiness, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新运单轨迹



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.updatePath.html

POST https://api.weixin.qq.com/cgi-bin/express/delivery/path/update?access_token=ACCESS_TOKEN
*/
func UpdatePath(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdatePath, bytes.NewReader(payload), "application/json;charset=utf-8")
}
