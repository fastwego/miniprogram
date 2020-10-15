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

var apiConfigWxopen = []ApiGroup{
	{
		Name:    `基础信息设置`,
		Package: `basic_info`,
		Apis: []Api{
			{
				Name:        "获取基本信息",
				Description: "调用本 API 可以获取小程序的基本信息。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Mini_Program_Information_Settings.html",
				FuncName:    "GetAccountBasicInfo",
			},
			{
				Name:        "设置服务器域名",
				Description: "授权给第三方的小程序，其服务器域名只可以为在第三方平台账号中配置的小程序服务器域名，当小程序通过第三方平台发布代码上线后，小程序原先自己配置的服务器域名将被删除，只保留第三方平台的域名，所以第三方平台在代替小程序发布代码之前，需要调用接口为小程序添加第三方平台自身的域名。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/modify_domain?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Server_Address_Configuration.html",
				FuncName:    "",
			},
			{
				Name:        "设置业务域名",
				Description: "授权给第三方的小程序，其业务域名只可以为在第三方平台账号中配置的小程序业务域名，当小程序通过第三方发布代码上线后，小程序原先自己配置的业务域名将被删除，只保留第三方平台的域名，所以第三方平台在代替小程序发布代码之前，需要调用接口为小程序添加业务域名。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/setwebviewdomain?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/setwebviewdomain.html",
				FuncName:    "SetWebviewDomain",
			},
			{
				Name:        "设置名称",
				Description: "调用本接口可以设置小程序名称，当名称没有命中关键词，则直接设置成功；当名称命中关键词，需提交证明材料，并需要审核。审核结果会向消息与事件接收 URL 进行事件推送。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/setnickname?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/setnickname.html",
				FuncName:    "SetNickname",
			},
			{
				Name:        "微信认证名称检测",
				Description: "调用本 API 可以检测微信认证的名称是否符合规则。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。POST https://api.weixin.qq.com/cgi-bin/wxverify/checkwxverifynickname?access_token=ACCESS_TOKEN",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxverify/checkwxverifynickname?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/wxverify_checknickname.html",
				FuncName:    "CheckWXVerifyNickname",
			},
			{
				Name:        "查询改名审核状态",
				Description: "调用设置名称接口，如果需要审核，会返回审核单 id（audit_id）,使用本接口可以查询改名审核状态。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/api_wxa_querynickname?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/api_wxa_querynickname.html",
				FuncName:    "QueryNickname",
			},
			{
				Name:        "修改头像",
				Description: "调用本接口可以修改小程序的头像。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/account/modifyheadimage?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/modifyheadimage.html",
				FuncName:    "ModifyHeadImage",
			},
			{
				Name:        "修改功能介绍",
				Description: "调用本接口可以修改功能介绍。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/account/modifysignature?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/modifysignature.html",
				FuncName:    "ModifySignature",
			},
			{
				Name:        "查询隐私设置",
				Description: "通过本接口可以查询小程序当前的隐私设置，即是否可被搜索。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/wxa/getwxasearchstatus?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/getwxasearchstatus.html",
				FuncName:    "GetWxaSearchStatus",
			},
			{
				Name:        "修改隐私设置",
				Description: "通过本接口修改小程序隐私设置，即修改是否可被搜索，可以先查询小程序当前的隐私设置再决定是否修改。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/changewxasearchstatus?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/changewxasearchstatus.html",
				FuncName:    "ChangeWxaSearchStatus",
			},
		},
	},
	{
		Name:    `类目管理`,
		Package: `category`,
		Apis: []Api{
			{
				Name:        "获取可以设置的所有类目",
				Description: "本接口可以获取该小程序允许设置的所有类目且仅支持获取一级类目和二级类目，注意不同主体所允许设置的类目不同。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/wxopen/getallcategories?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/getallcategories.html",
				FuncName:    "GetAllCategories",
			},
			{
				Name:        "获取已设置的所有类目",
				Description: "使用本接口获取小程序已设置的所有类目。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/wxopen/getcategory?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/getcategory.html",
				FuncName:    "GetCategory",
			},
			{
				Name:        "添加类目",
				Description: "调用本接口可以给小程序添加类目，添加的类目需要在所有可设置的类目列表中。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxopen/addcategory?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/addcategory.html",
				FuncName:    "AddCategory",
			},
			{
				Name:        "删除类目",
				Description: "调用本接口可以删除小程序的指定类目。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxopen/deletecategory?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/deletecategory.html",
				FuncName:    "DeleteCategory",
			},
			{
				Name:        "修改类目资质信息",
				Description: "通过获取已设置的类目列表接口 可以获取当前小程序已设置的类目列表。如果某一下类目审核不通过需要补充或者修改资质信息，可以调用本接口进行处理。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxopen/modifycategory?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/modifycategory.html",
				FuncName:    "ModifyCategory",
			},
			{
				Name:        "获取审核时可填写的类目信息",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxa/get_category?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/category/get_category.html",
				FuncName:    "WxaCategory",
			},
		},
	},
	{
		Name:    `扫码关注组件`,
		Package: `subscribe_component`,
		Apis: []Api{

			{
				Name:        "获取展示的公众号信息",
				Description: "使用本接口可以获取扫码关注组件所展示的公众号信息",
				Request:     "GET https://api.weixin.qq.com/wxa/getshowwxaitem?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_component/getshowwxaitem.html",
				FuncName:    "GetShowWxaItem",
			},
			{
				Name:        "获取可以用来设置的公众号列表",
				Description: "通过本接口可以获取扫码关注组件允许展示的公众号列表",
				Request:     "GET https://api.weixin.qq.com/wxa/getwxamplinkforshow?page=0&num=20&access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_component/getwxamplinkforshow.html",
				FuncName:    "GetWxaMplinkForShow",
				GetParams: []Param{
					{Name: `page`, Type: `string`},
					{Name: `num`, Type: `string`},
				},
			},
			{
				Name:        "设置展示的公众号信息",
				Description: "使用本接口可以设置扫码关注组件所展示的公众号信息",
				Request:     "POST https://api.weixin.qq.com/wxa/updateshowwxaitem?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_component/updateshowwxaitem.html",
				FuncName:    "UpdateShowWxaItem",
			},
		},
	},
	{
		Name:    `普通链接二维码与小程序码`,
		Package: `qrcode`,
		Apis: []Api{

			{
				Name:        "获取已设置的二维码规则",
				Description: "通过本接口可以获取已设置的普通链接二维码规则。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpget?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/qrcodejumpget.html",
				FuncName:    "QRCodeJumpGet",
			},
			{
				Name:        "获取校验文件名称及内容",
				Description: "通过本接口下载随机校验文件，并将文件上传至服务器指定位置的目录下，方可通过所属权校验。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpdownload?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/qrcodejumpdownload.html",
				FuncName:    "QRCodeJumpDownload",
			},
			{
				Name:        "增加或修改二维码规则",
				Description: "通过本接口可以增加或修改普通链接二维码规则.使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpadd?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/qrcodejumpadd.html",
				FuncName:    "QRCodeJumpAdd",
			},
			{
				Name:        "发布已设置的二维码规则",
				Description: "需要先添加二维码规则，然后调用本接口将二维码规则发布生效，发布后现网用户扫码命中改规则的普通链接二维码时将调整到正式版小程序指定的页面。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumppublish?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/qrcodejumppublish.html",
				FuncName:    "QRCodeJumpPublish",
			},
			{
				Name:        "删除已设置的二维码规则",
				Description: "该接口用于删除已设置的二维码规则。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpdelete?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/qrcodejumpdelete.html",
				FuncName:    "QRCodeJumpDelete",
			},
			{
				Name:        "将二维码长链接转成短链接",
				Description: "调用本 API 可以将一条长链接转成短链接。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/shorturl?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/shorturl.html",
				FuncName:    "",
			},
			{
				Name:        "获取unlimit小程序码",
				Description: "调用本 API 可以获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。 使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/getwxacodeunlimit.html",
				FuncName:    "GetWxaCodeUnLimit",
			},
			{
				Name:        "获取小程序码",
				Description: "调用本 API 可以获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/getwxacode.html",
				FuncName:    "GetWxaCode",
			},
			{
				Name:        "获取小程序二维码",
				Description: "调用本 API 可以获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/qrcode/createwxaqrcode.html",
				FuncName:    "CreateWxaQRCode",
			},
		},
	},
	{
		Name:    `成员管理`,
		Package: `member_admin`,
		Apis: []Api{

			{
				Name:        "绑定微信用户为体验者",
				Description: "第三方平台在帮助旗下授权的小程序提交代码审核之前，可先让小程序运营者体验，体验之前需要将运营者的个人微信号添加到该小程序的体验者名单中。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/bind_tester?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Admin.html",
				FuncName:    "",
			},
			{
				Name:        "解除绑定体验者",
				Description: "调用本接口可以将特定微信用户从小程序的体验者列表中解绑。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/unbind_tester?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/unbind_tester.html",
				FuncName:    "",
			},
			{
				Name:        "获取体验者列表",
				Description: "通过本接口可以获取小程序所有已绑定的体验者列表。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/memberauth?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/memberauth.html",
				FuncName:    "MemberAuth",
			},
		},
	}, {
		Name:    `代码模板库设置`,
		Package: `code_template_library_management`,
		Apis: []Api{

			{
				Name:        "获取代码草稿列表",
				Description: "通过本接口，可以获取草稿箱中所有的草稿（临时代码模板）；草稿是由第三方平台的开发小程序在使用微信开发者工具上传的，点击查ext详情使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/wxa/gettemplatedraftlist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatedraftlist.html",
				FuncName:    "GetTemplateDraftList",
			},
			{
				Name:        "将草稿添加到代码模板库",
				Description: "可以通过获取草稿箱中所有的草稿得到草稿 ID；调用本接口可以将临时草稿选为持久的代码模板。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/addtotemplate?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/addtotemplate.html",
				FuncName:    "AddToTemplate",
			},
			{
				Name:        "获取代码模板列表",
				Description: "第三方平台运营者可以登录 open.weixin.qq.com 或者通过将草稿箱的草稿选为代码模板接口，将草稿箱中的某个代码版本添加到代码模板库中使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/wxa/gettemplatelist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatelist.html",
				FuncName:    "GetTemplateList",
			},
			{
				Name:        "删除指定代码模板",
				Description: "因为代码模板库的模板数量是有上限的，当达到上限或者有某个模板不再需要时，可以调用本接口删除指定的代码模板。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/deletetemplate?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/deletetemplate.html",
				FuncName:    "DeleteTemplate",
			},
		},
	},
	{
		Name:    `代码管理`,
		Package: `code_management`,
		Apis: []Api{

			{
				Name:        "上传小程序代码",
				Description: "第三方平台需要先将草稿添加到代码模板库，或者从代码模板库中选取某个代码模板，得到对应的模板 id（template_id）； 然后调用本接口可以为已授权的小程序上传代码。",
				Request:     "POST https://api.weixin.qq.com/wxa/commit?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/commit.html",
				FuncName:    "",
			},
			{
				Name:        "获取已上传的代码的页面列表",
				Description: "通过本接口可以获取由第三方平台上传小程序代码的页面列表；用于提交审核的审核项 的 address 参数使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/wxa/get_page?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_page.html",
				FuncName:    "",
			},
			{
				Name:        "获取体验版二维码",
				Description: "调用本接口可以获取小程序的体验版二维码。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/wxa/get_qrcode?access_token=ACCESS_TOKEN&path=page%2Findex%3Faction%3D1",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_qrcode.html",
				FuncName:    "",
				GetParams: []Param{
					{Name: `path`, Type: `string`},
				},
			},
			{
				Name:        "提交审核",
				Description: "在调用上传代码接口为小程序上传代码后，可以调用本接口，将上传的代码提交审核。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/submit_audit?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/submit_audit.html",
				FuncName:    "",
			},
			{
				Name:        "查询指定发布审核单的审核状态",
				Description: "提交审核后，调用本接口可以查询指定发布审核单的审核状态。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/get_auditstatus?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_auditstatus.html",
				FuncName:    "",
			},
			{
				Name:        "查询最新一次提交的审核状态",
				Description: "调用本接口可以查询最新一次提审单的审核状态。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/wxa/get_latest_auditstatus?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_latest_auditstatus.html",
				FuncName:    "",
			},
			{
				Name:        "小程序审核撤回",
				Description: "调用本接口可以撤回当前的代码审核单。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/wxa/undocodeaudit?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/undocodeaudit.html",
				FuncName:    "UndoCodeAudit",
			},
			{
				Name:        "发布已通过审核的小程序",
				Description: "调用本接口可以发布最后一个审核通过的小程序代码版本。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/release?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/release.html",
				FuncName:    "",
			},
			{
				Name:        "版本回退",
				Description: "调用本接口可以将小程序的线上版本进行回退。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/wxa/revertcoderelease?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/revertcoderelease.html",
				FuncName:    "RevertCodeRelease",
			},
			{
				Name:        "分阶段发布",
				Description: "发布小程序接口 是全量发布，会影响到现网的所有用户。而本接口是创建一个灰度发布的计划，可以控制发布的节奏，避免一上线就影响到所有的用户。可以多次调用本次接口，将灰度的比例（gray_percentage）逐渐增大使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/grayrelease?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/grayrelease.html",
				FuncName:    "GrayRelease",
			},
			{
				Name:        "查询当前分阶段发布详情",
				Description: "该接口用于查询当前分阶段发布详情。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/wxa/getgrayreleaseplan?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/getgrayreleaseplan.html",
				FuncName:    "GetGrayReleasePlan",
			},
			{
				Name:        "取消分阶段发布",
				Description: "在小程序分阶段发布期间，可以随时调用本接口取消分阶段发布。取消分阶段发布后，受影响的微信用户（即被灰度升级的微信用户）的小程序版本将回退到分阶段发布前的版本",
				Request:     "GET https://api.weixin.qq.com/wxa/revertgrayrelease?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/revertgrayrelease.html",
				FuncName:    "RevertGrayRelease",
			},
			{
				Name:        "修改小程序线上代码的可见状态",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/change_visitstatus?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/change_visitstatus.html",
				FuncName:    "",
			},
			{
				Name:        "查询当前设置的最低基础库版本及各版本用户占比",
				Description: "调用本接口可以查询小程序当前设置的最低基础库版本，以及小程序在各个基础库版本的用户占比。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxopen/getweappsupportversion?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/getweappsupportversion.html",
				FuncName:    "GetWeAppSupportVersion",
			},
			{
				Name:        "设置最低基础库版本",
				Description: "调用本接口可以设置小程序的最低基础库支持版本，可以先查询当前小程序在各个基础库的用户占比来辅助进行决策使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/wxopen/setweappsupportversion?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/setweappsupportversion.html",
				FuncName:    "SetWeAppSupportVersion",
			},
			{
				Name:        "查询服务商的当月提审限额（quota）和加急次数",
				Description: "服务商可以调用该接口，查询当月平台分配的提审限额和剩余可提审次数，以及当月分配的审核加急次数和剩余加急次数。（所有旗下小程序共用该额度）使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "GET https://api.weixin.qq.com/wxa/queryquota?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/query_quota.html",
				FuncName:    "QueryQuota",
			},
			{
				Name:        "加急审核申请",
				Description: "有加急次数的第三方可以通过该接口，对已经提审的小程序进行加急操作，加急后的小程序预计2-12小时内审完。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。",
				Request:     "POST https://api.weixin.qq.com/wxa/speedupaudit?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/speedup_audit.html",
				FuncName:    "SpeedupAudit",
			},
		},
	},
}
