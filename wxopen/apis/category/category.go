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

// Package category 类目管理
package category

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiGetAllCategories = "/cgi-bin/wxopen/getallcategories"
	apiGetCategory      = "/cgi-bin/wxopen/getcategory"
	apiAddCategory      = "/cgi-bin/wxopen/addcategory"
	apiDeleteCategory   = "/cgi-bin/wxopen/deletecategory"
	apiModifyCategory   = "/cgi-bin/wxopen/modifycategory"
	apiWxaCategory      = "/wxa/get_category"
)

/*
获取可以设置的所有类目

本接口可以获取该小程序允许设置的所有类目且仅支持获取一级类目和二级类目，注意不同主体所允许设置的类目不同。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/getallcategories.html

GET https://api.weixin.qq.com/cgi-bin/wxopen/getallcategories?access_token=ACCESS_TOKEN
*/
func GetAllCategories(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetAllCategories)
}

/*
获取已设置的所有类目

使用本接口获取小程序已设置的所有类目。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/getcategory.html

GET https://api.weixin.qq.com/cgi-bin/wxopen/getcategory?access_token=ACCESS_TOKEN
*/
func GetCategory(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetCategory)
}

/*
添加类目

调用本接口可以给小程序添加类目，添加的类目需要在所有可设置的类目列表中。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/addcategory.html

POST https://api.weixin.qq.com/cgi-bin/wxopen/addcategory?access_token=ACCESS_TOKEN
*/
func AddCategory(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddCategory, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除类目

调用本接口可以删除小程序的指定类目。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/deletecategory.html

POST https://api.weixin.qq.com/cgi-bin/wxopen/deletecategory?access_token=TOKEN
*/
func DeleteCategory(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDeleteCategory, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改类目资质信息

通过获取已设置的类目列表接口 可以获取当前小程序已设置的类目列表。如果某一下类目审核不通过需要补充或者修改资质信息，可以调用本接口进行处理。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/modifycategory.html

POST https://api.weixin.qq.com/cgi-bin/wxopen/modifycategory?access_token=ACCESS_TOKEN
*/
func ModifyCategory(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiModifyCategory, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取审核时可填写的类目信息



See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/get_category.html

GET https://api.weixin.qq.com/wxa/get_category?access_token=ACCESS_TOKEN
*/
func WxaCategory(ctx *miniprogram.Miniprogram) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiWxaCategory)
}
