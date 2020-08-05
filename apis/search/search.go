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

// Package search 小程序搜索
package search

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiImageSearch = "/wxa/imagesearch"
	apiSiteSearch  = "/wxa/sitesearch"
	apiSubmitPages = "/wxa/search/wxaapi_submitpages"
)

/*
本接口提供基于小程序的站内搜商品图片搜索能力



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.imageSearch.html

POST https://api.weixin.qq.com/wxa/imagesearch?access_token=TOKEN
*/
func ImageSearch(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiImageSearch, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
小程序内部搜索API提供针对页面的查询能力，小程序开发者输入搜索词后，将返回自身小程序和搜索词相关的页面。因此，利用该接口，开发者可以查看指定内容的页面被微信平台的收录情况；同时，该接口也可供开发者在小程序内应用，给小程序用户提供搜索能力。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.siteSearch.html

POST https://api.weixin.qq.com/wxa/sitesearch?access_token=TOKEN
*/
func SiteSearch(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSiteSearch, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
小程序开发者可以通过本接口提交小程序页面url及参数信息(不要推送webview页面)，让微信可以更及时的收录到小程序的页面信息，开发者提交的页面信息将可能被用于小程序搜索结果展示。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.submitPages.html

POST https://api.weixin.qq.com/wxa/search/wxaapi_submitpages?access_token=TOKEN
*/
func SubmitPages(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSubmitPages, bytes.NewReader(payload), "application/json;charset=utf-8")
}
