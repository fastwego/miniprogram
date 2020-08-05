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

// Package datacube 数据分析
package datacube

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiGetDailyRetain       = "/datacube/getweanalysisappiddailyretaininfo"
	apiGetMonthlyRetain     = "/datacube/getweanalysisappidmonthlyretaininfo"
	apiGetWeeklyRetain      = "/datacube/getweanalysisappidweeklyretaininfo"
	apiGetDailySummary      = "/datacube/getweanalysisappiddailysummarytrend"
	apiGetDailyVisitTrend   = "/datacube/getweanalysisappiddailyvisittrend"
	apiGetMonthlyVisitTrend = "/datacube/getweanalysisappidmonthlyvisittrend"
	apiGetWeeklyVisitTrend  = "/datacube/getweanalysisappidweeklyvisittrend"
	apiGetUserPortrait      = "/datacube/getweanalysisappiduserportrait"
	apiGetVisitDistribution = "/datacube/getweanalysisappidvisitdistribution"
	apiGetVisitPage         = "/datacube/getweanalysisappidvisitpage"
)

/*
获取用户访问小程序日留存



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html

POST https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=ACCESS_TOKEN
*/
func GetDailyRetain(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetDailyRetain, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户访问小程序月留存



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html

POST https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token=ACCESS_TOKEN
*/
func GetMonthlyRetain(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetMonthlyRetain, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户访问小程序周留存



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html

POST https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token=ACCESS_TOKEN
*/
func GetWeeklyRetain(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetWeeklyRetain, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户访问小程序数据概况



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html

POST https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token=ACCESS_TOKEN
*/
func GetDailySummary(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetDailySummary, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户访问小程序数据日趋势



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html

POST https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token=ACCESS_TOKEN
*/
func GetDailyVisitTrend(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetDailyVisitTrend, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户访问小程序数据月趋势(能查询到的最新数据为上一个自然月的数据)



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html

POST https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend?access_token=ACCESS_TOKEN
*/
func GetMonthlyVisitTrend(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetMonthlyVisitTrend, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户访问小程序数据周趋势



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getWeeklyVisitTrend.html

POST https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend?access_token=ACCESS_TOKEN
*/
func GetWeeklyVisitTrend(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetWeeklyVisitTrend, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取小程序新增或活跃用户的画像分布数据。时间范围支持昨天、最近7天、最近30天。其中，新增用户数为时间范围内首次访问小程序的去重用户数，活跃用户数为时间范围内访问过小程序的去重用户数。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html

POST https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait?access_token=ACCESS_TOKEN
*/
func GetUserPortrait(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserPortrait, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户小程序访问分布数据



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitDistribution.html

POST https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution?access_token=ACCESS_TOKEN
*/
func GetVisitDistribution(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetVisitDistribution, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
访问页面。目前只提供按 page_visit_pv 排序的 top200。



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitPage.html

POST https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token=ACCESS_TOKEN
*/
func GetVisitPage(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetVisitPage, bytes.NewReader(payload), "application/json;charset=utf-8")
}
