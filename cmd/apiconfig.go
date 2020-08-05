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

package main

type Param struct {
	Name string
	Type string
}

type Api struct {
	Name        string
	Description string
	Request     string
	See         string
	FuncName    string
	GetParams   []Param
}

type ApiGroup struct {
	Name    string
	Apis    []Api
	Package string
}

var apiConfig = []ApiGroup{
	{
		Name:    `用户`,
		Package: `user`,
		Apis: []Api{
			{
				Name:        "登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。更多使用方法详见 小程序登录。",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html",
				FuncName:    "Code2Session",
				GetParams: []Param{
					{Name: `appid`, Type: `string`},
					{Name: `secret`, Type: `string`},
					{Name: `js_code`, Type: `string`},
					{Name: `grant_type`, Type: `string`},
				},
			},
			{
				Name:        "用户支付完成后，获取该用户的 UnionId，无需用户授权。本接口支持第三方平台代理查询。",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxa/getpaidunionid?access_token=ACCESS_TOKEN&openid=OPENID",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html",
				FuncName:    "GetPaidUnionId",
				GetParams: []Param{
					{Name: `openid`, Type: `string`},
				},
			},
		},
	},
	{
		Name:    `数据分析`,
		Package: `datacube`,
		Apis: []Api{

			{
				Name:        "获取用户访问小程序日留存",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html",
				FuncName:    "GetDailyRetain",
			},
			{
				Name:        "获取用户访问小程序月留存",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html",
				FuncName:    "GetMonthlyRetain",
			},
			{
				Name:        "获取用户访问小程序周留存",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html",
				FuncName:    "GetWeeklyRetain",
			},
			{
				Name:        "获取用户访问小程序数据概况",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html",
				FuncName:    "GetDailySummary",
			},
			{
				Name:        "获取用户访问小程序数据日趋势",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html",
				FuncName:    "GetDailyVisitTrend",
			},
			{
				Name:        "获取用户访问小程序数据月趋势(能查询到的最新数据为上一个自然月的数据)",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html",
				FuncName:    "GetMonthlyVisitTrend",
			},
			{
				Name:        "获取用户访问小程序数据周趋势",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getWeeklyVisitTrend.html",
				FuncName:    "GetWeeklyVisitTrend",
			},
			{
				Name:        "获取小程序新增或活跃用户的画像分布数据。时间范围支持昨天、最近7天、最近30天。其中，新增用户数为时间范围内首次访问小程序的去重用户数，活跃用户数为时间范围内访问过小程序的去重用户数。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html",
				FuncName:    "GetUserPortrait",
			},
			{
				Name:        "获取用户小程序访问分布数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitDistribution.html",
				FuncName:    "GetVisitDistribution",
			},
			{
				Name:        "访问页面。目前只提供按 page_visit_pv 排序的 top200。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitPage.html",
				FuncName:    "GetVisitPage",
			},
		},
	},
	{
		Name:    `客服消息`,
		Package: `customer_message`,
		Apis: []Api{

			{
				Name:        "获取客服消息内的临时素材。即下载临时的多媒体文件。目前小程序仅支持下载图片文件。",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.getTempMedia.html",
				FuncName:    "GetTempMedia",
				GetParams: []Param{
					{Name: `media_id`, Type: `string`},
				},
			},
			{
				Name:        "发送客服消息给用户。详细规则见 发送客服消息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html",
				FuncName:    "Send",
			},
			{
				Name:        "下发客服当前输入状态给用户。详见 客服消息输入状态",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html",
				FuncName:    "SetTyping",
			},
			{
				Name:        "把媒体文件上传到微信服务器。目前仅支持图片。用于发送客服消息或被动回复用户消息。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html",
				FuncName:    "UploadTempMedia",
				GetParams: []Param{
					{Name: `type`, Type: `string`},
				},
			},
		},
	},
	{
		Name:    `统一服务消息`,
		Package: `uniform_message`,
		Apis: []Api{
			{
				Name:        "下发小程序和公众号统一的服务消息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html",
				FuncName:    "Send",
			},
		},
	},
	{
		Name:    `动态消息`,
		Package: `updatable_message`,
		Apis: []Api{
			{
				Name:        "创建被分享动态消息的 activity_id。详见动态消息。",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/message/wxopen/activityid/create?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.createActivityId.html",
				FuncName:    "CreateActivityId",
			},
			{
				Name:        "修改被分享的动态消息。详见动态消息。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/wxopen/updatablemsg/send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.setUpdatableMsg.html",
				FuncName:    "SetUpdatableMsg",
			},
		},
	},
	{
		Name:    `插件管理`,
		Package: `plugin`,
		Apis: []Api{

			{
				Name:        "向插件开发者发起使用插件的申请/查询已添加的插件/删除已添加的插件",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/plugin?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.applyPlugin.html",
				FuncName:    "ApplyPlugin",
			},
			{
				Name:        "获取当前所有插件使用方/修改插件使用申请的状态",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/devplugin?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginDevApplyList.html",
				FuncName:    "GetPluginDevApplyList",
			},
		},
	},
	{
		Name:    `附近的小程序`,
		Package: `nearbypoi`,
		Apis: []Api{

			{
				Name:        "添加地点",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/addnearbypoi?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.add.html",
				FuncName:    "Add",
			},
			{
				Name:        "删除地点",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/delnearbypoi?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.delete.html",
				FuncName:    "Delete",
			},
			{
				Name:        "查看地点列表",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxa/getnearbypoilist?page=1&page_rows=20&access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.getList.html",
				FuncName:    "GetList",
				GetParams: []Param{
					{Name: `page`, Type: `string`},
					{Name: `page_rows`, Type: `string`},
				},
			},
			{
				Name:        "展示/取消展示附近小程序",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/setnearbypoishowstatus?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.setShowStatus.html",
				FuncName:    "SetShowStatus",
			},
		},
	},
	{
		Name:    `小程序码`,
		Package: `wxacode`,
		Apis: []Api{
			{
				Name:        "获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html",
				FuncName:    "CreateQRCode",
			},
			{
				Name:        "获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html",
				FuncName:    "Get",
			},
			{
				Name:        "获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。 更多用法详见 获取二维码。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html",
				FuncName:    "GetUnlimited",
			},
		},
	},
	{
		Name:    `内容安全`,
		Package: `security`,
		Apis: []Api{

			{
				Name:        "校验一张图片是否含有违法违规内容。",
				Description: "",
				Request:     "POST(@media) https://api.weixin.qq.com/wxa/img_sec_check?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html",
				FuncName:    "ImgSecCheck",
			},
			{
				Name:        "异步校验图片/音频是否含有违法违规内容。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/media_check_async?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html",
				FuncName:    "MediaCheckAsync",
			},
			{
				Name:        "检查一段文本是否含有违法违规内容。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/msg_sec_check?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.msgSecCheck.html",
				FuncName:    "MsgSecCheck",
			},
		},
	},
	{
		Name:    `物流助手/小程序使用`,
		Package: `logistics/mini`,
		Apis: []Api{

			{
				Name:        "批量获取运单数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/business/order/batchget?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.batchGetOrder.html",
				FuncName:    "BatchGetOrder",
			},
			{
				Name:        "生成运单",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/business/order/add?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.addOrder.html",
				FuncName:    "AddOrder",
			},
			{
				Name:        "绑定、解绑物流账号",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/business/account/bind?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.bindAccount.html",
				FuncName:    "BindAccount",
			},
			{
				Name:        "取消运单",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/business/order/cancel?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.cancelOrder.html",
				FuncName:    "CancelOrder",
			},
			{
				Name:        "获取所有绑定的物流账号",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/express/business/account/getall?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getAllAccount.html",
				FuncName:    "GetAllAccount",
			},
			{
				Name:        "获取支持的快递公司列表",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/express/business/delivery/getall?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getAllDelivery.html",
				FuncName:    "GetAllDelivery",
			},
			{
				Name:        "获取运单数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/business/order/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getOrder.html",
				FuncName:    "GetOrder",
			},
			{
				Name:        "查询运单轨迹",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/business/path/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getPath.html",
				FuncName:    "GetPath",
			},
			{
				Name:        "获取打印员。若需要使用微信打单 PC 软件，才需要调用。",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/express/business/printer/getall?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getPrinter.html",
				FuncName:    "GetPrinter",
			},
			{
				Name:        "获取电子面单余额。仅在使用加盟类快递公司时，才可以调用。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/business/quota/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getQuota.html",
				FuncName:    "GetQuota",
			},
			{
				Name:        "模拟快递公司更新订单状态, 该接口只能用户测试",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/business/test_update_order?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.testUpdateOrder.html",
				FuncName:    "TestUpdateOrder",
			},
			{
				Name:        "配置面单打印员，可以设置多个，若需要使用微信打单 PC 软件，才需要调用。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/business/printer/update?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.updatePrinter.html",
				FuncName:    "UpdatePrinter",
			},
		},
	},
	{
		Name:    `物流助手/运力方`,
		Package: `logistics/provider`,
		Apis: []Api{

			{
				Name:        "获取面单联系人信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/delivery/contact/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.getContact.html",
				FuncName:    "GetContact",
			},
			{
				Name:        "预览面单模板。用于调试面单模板使用。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/delivery/template/preview?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.previewTemplate.html",
				FuncName:    "PreviewTemplate",
			},
			{
				Name:        "更新商户审核结果",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/delivery/service/business/update?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.updateBusiness.html",
				FuncName:    "UpdateBusiness",
			},
			{
				Name:        "更新运单轨迹",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/delivery/path/update?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.updatePath.html",
				FuncName:    "UpdatePath",
			},
		},
	},
	{
		Name:    `广告`,
		Package: `ad`,
		Apis:    []Api{},
	},
	{
		Name:    `图像处理`,
		Package: `img`,
		Apis: []Api{

			{
				Name:        "本接口提供基于小程序的图片智能裁剪能力。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/img/aicrop?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.aiCrop.html",
				FuncName:    "AiCrop",
				GetParams: []Param{
					{Name: `img_url`, Type: `string`},
				},
			},
			{
				Name:        "本接口提供基于小程序的条码/二维码识别的API。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/img/qrcode?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.scanQRCode.html",
				FuncName:    "ScanQRCode",
				GetParams: []Param{
					{Name: `img_url`, Type: `string`},
				},
			},
			{
				Name:        "本接口提供基于小程序的图片高清化能力。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/img/superresolution?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.superresolution.html",
				FuncName:    "Superresolution",
				GetParams: []Param{
					{Name: `img_url`, Type: `string`},
				},
			},
		},
	},
	{
		Name:    `即时配送/小程序使用`,
		Package: `delivery/mini`,
		Apis: []Api{

			{
				Name:        "异常件退回商家商家确认收货接口",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/confirm_return?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.abnormalConfirm.html",
				FuncName:    "AbnormalConfirm",
			},
			{
				Name:        "下配送单接口",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/add?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addOrder.html",
				FuncName:    "AddOrder",
			},
			{
				Name:        "可以对待接单状态的订单增加小费。需要注意：订单的小费，以最新一次加小费动作的金额为准，故下一次增加小费额必须大于上一次小费额",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/addtips?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addTip.html",
				FuncName:    "AddTip",
			},
			{
				Name:        "第三方代商户发起绑定配送公司帐号的请求",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/shop/add?access_token=COMPNENT_ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.bindAccount.html",
				FuncName:    "BindAccount",
			},
			{
				Name:        "取消配送单接口",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/cancel?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.cancelOrder.html",
				FuncName:    "CancelOrder",
			},
			{
				Name:        "获取已支持的配送公司列表接口",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/delivery/getall?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getAllImmeDelivery.html",
				FuncName:    "GetAllImmeDelivery",
			},
			{
				Name:        "拉取已绑定账号",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/shop/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getBindAccount.html",
				FuncName:    "GetBindAccount",
			},
			{
				Name:        "拉取配送单信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getOrder.html",
				FuncName:    "GetOrder",
			},
			{
				Name:        "模拟配送公司更新配送单状态, 该接口只用于沙盒环境，即订单并没有真实流转到运力方",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/test_update_order?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.mockUpdateOrder.html",
				FuncName:    "MockUpdateOrder",
			},
			{
				Name:        "第三方代商户发起开通即时配送权限",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/open?access_token=COMPNENT_ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.openDelivery.html",
				FuncName:    "OpenDelivery",
			},
			{
				Name:        "预下配送单接口",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/pre_add?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.preAddOrder.html",
				FuncName:    "PreAddOrder",
			},
			{
				Name:        "预取消配送单接口",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/precancel?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.preCancelOrder.html",
				FuncName:    "PreCancelOrder",
			},
			{
				Name:        "模拟配送公司更新配送单状态, 该接口用于测试账户下的单，将请求转发到运力测试环境",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/realmock_update_order?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.realMockUpdateOrder.html",
				FuncName:    "RealMockUpdateOrder",
			},
			{
				Name:        "重新下单",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/business/order/readd?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.reOrder.html",
				FuncName:    "ReOrder",
			},
		},
	}, {
		Name:    `即时配送/运力方`,
		Package: `delivery/provider`,
		Apis: []Api{
			{
				Name:        "配送公司更新配送单状态",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/express/local/delivery/update_order?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.updateOrder.html",
				FuncName:    "UpdateOrder",
			},
		},
	},
	{
		Name:    `OCR`,
		Package: `ocr`,
		Apis: []Api{

			{
				Name:        "本接口提供基于小程序的银行卡 OCR 识别",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/bankcard?type=MODE&img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.bankcard.html",
				FuncName:    "Bankcard",
				GetParams: []Param{
					{Name: `type`, Type: `string`},
					{Name: `img_url`, Type: `string`},
				},
			},
			{
				Name:        "本接口提供基于小程序的营业执照 OCR 识别",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/bizlicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.businessLicense.html",
				FuncName:    "BusinessLicense",
				GetParams: []Param{
					{Name: `img_url`, Type: `string`},
				},
			},
			{
				Name:        "本接口提供基于小程序的驾驶证 OCR 识别",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/drivinglicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.driverLicense.html",
				FuncName:    "DriverLicense",
				GetParams: []Param{
					{Name: `img_url`, Type: `string`},
				},
			},
			{
				Name:        "本接口提供基于小程序的身份证 OCR 识别",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/idcard?type=MODE&img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.idcard.html",
				FuncName:    "Idcard",
				GetParams: []Param{
					{Name: `type`, Type: `string`},
					{Name: `img_url`, Type: `string`},
				},
			},
			{
				Name:        "本接口提供基于小程序的通用印刷体 OCR 识别",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/comm?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.printedText.html",
				FuncName:    "PrintedText",
				GetParams: []Param{
					{Name: `img_url`, Type: `string`},
				},
			},
			{
				Name:        "本接口提供基于小程序的行驶证 OCR 识别",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/driving?type=MODE&img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.vehicleLicense.html",
				FuncName:    "VehicleLicense",
				GetParams: []Param{
					{Name: `img_url`, Type: `string`},
					{Name: `type`, Type: `string`},
				},
			},
		},
	},
	{
		Name:    `运维中心`,
		Package: `operation`,
		Apis: []Api{

			{
				Name:        "获取用户反馈列表",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxaapi/feedback/list?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getFeedback.html",
				FuncName:    "GetFeedback",
			},
			{
				Name:        "错误查询",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxaapi/log/jserr_search?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getJsErrSearch.html",
				FuncName:    "GetJsErrSearch",
			},
			{
				Name:        "性能监控",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxaapi/log/get_performance?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getPerformance.html",
				FuncName:    "GetPerformance",
			},
			{
				Name:        "获取访问来源",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxaapi/log/get_scene?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getSceneList.html",
				FuncName:    "GetSceneList",
			},
			{
				Name:        "获取客户端版本",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxaapi/log/get_client_version?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getVersionList.html",
				FuncName:    "GetVersionList",
			},
			{
				Name:        "实时日志查询",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxaapi/userlog/userlog_search?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.realtimelogSearch.html",
				FuncName:    "RealtimelogSearch",
			},
		},
	},
	{
		Name:    `小程序搜索`,
		Package: `search`,
		Apis: []Api{

			{
				Name:        "本接口提供基于小程序的站内搜商品图片搜索能力",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/imagesearch?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.imageSearch.html",
				FuncName:    "ImageSearch",
			},
			{
				Name:        "小程序内部搜索API提供针对页面的查询能力，小程序开发者输入搜索词后，将返回自身小程序和搜索词相关的页面。因此，利用该接口，开发者可以查看指定内容的页面被微信平台的收录情况；同时，该接口也可供开发者在小程序内应用，给小程序用户提供搜索能力。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/sitesearch?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.siteSearch.html",
				FuncName:    "SiteSearch",
			},
			{
				Name:        "小程序开发者可以通过本接口提交小程序页面url及参数信息(不要推送webview页面)，让微信可以更及时的收录到小程序的页面信息，开发者提交的页面信息将可能被用于小程序搜索结果展示。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/search/wxaapi_submitpages?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.submitPages.html",
				FuncName:    "SubmitPages",
			},
		},
	},
	{
		Name:    `服务市场`,
		Package: `service_market`,
		Apis: []Api{
			{
				Name:        "调用服务平台提供的服务",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/servicemarket?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/service-market/serviceMarket.invokeService.html",
				FuncName:    "InvokeService",
			},
		},
	},
	{
		Name:    `生物认证`,
		Package: `verify`,
		Apis: []Api{

			{
				Name:        "SOTER 生物认证秘钥签名验证",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/soter/verify_signature?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/soter/soter.verifySignature.html",
				FuncName:    "VerifySignature",
			},
		},
	},
	{
		Name:    `订阅消息`,
		Package: `subscribe`,
		Apis: []Api{

			{
				Name:        "组合模板并添加至帐号下的个人模板库",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.addTemplate.html",
				FuncName:    "AddTemplate",
			},
			{
				Name:        "删除帐号下的个人模板",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.deleteTemplate.html",
				FuncName:    "DeleteTemplate",
			},
			{
				Name:        "获取小程序账号的类目",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxaapi/newtmpl/getcategory?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getCategory.html",
				FuncName:    "GetCategory",
			},
			{
				Name:        "获取模板标题下的关键词列表",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateKeyWordsById.html",
				FuncName:    "GetPubTemplateKeyWordsById",
			},
			{
				Name:        "获取帐号所属类目下的公共模板标题",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateTitleList.html",
				FuncName:    "GetPubTemplateTitleList",
			},
			{
				Name:        "获取当前帐号下的个人模板列表",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html",
				FuncName:    "GetTemplateList",
			},
			{
				Name:        "发送订阅消息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html",
				FuncName:    "Send",
			},
		},
	},
}
