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

var apiConfigMinigame = []ApiGroup{
	{
		Name:    `虚拟支付`,
		Package: `pay`,
		Apis: []Api{
			{
				Name:        "取消订单。开通了虚拟支付的小游戏，若扣除游戏币的订单号在有效时间内，可以通过本接口取消该笔扣除游戏币的订单",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/midas/cancelpay?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.cancelPay.html",
				FuncName:    "CancelPay",
			},
			{
				Name:        "获取游戏币余额。开通了虚拟支付的小游戏，可以通过本接口查看某个用户的游戏币余额",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/midas/getbalance?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.getBalance.html",
				FuncName:    "GetBalance",
			},
			{
				Name:        "扣除游戏币。开通了虚拟支付的小游戏，可以通过本接口扣除某个用户的游戏币。 由于可能存在接口调用超时或返回系统失败，但是游戏币实际已经扣除的情况，所以当该接口返回系统失败时，可以用相同的bill_no再次调用本接口，直到返回非系统失败为止，不会重复扣款，也可以调用取消支付接口取消本次扣款。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/midas/pay?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.pay.html",
				FuncName:    "Pay",
			},
			{
				Name:        "给用户赠送游戏币。开通了虚拟支付的小游戏，可以通过该接口赠送游戏币给某个用户。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/midas/present?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/midas-payment/midas.present.html",
				FuncName:    "Present",
			},
		},
	},
	{
		Name:    `用户`,
		Package: `user`,
		Apis: []Api{

			{
				Name:        "校验服务器所保存的登录态 session_key 是否合法。为了保持 session_key 私密性，接口不明文传输 session_key，而是通过校验登录态签名完成。",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxa/checksession?access_token=ACCESS_TOKEN&signature=SIGNATURE&openid=OPENID&sig_method=SIG_METHOD",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/login/auth.checkSessionKey.html",
				FuncName:    "CheckSessionKey",
				GetParams: []Param{
					{Name: `signature`, Type: `string`},
					{Name: `openid`, Type: `string`},
					{Name: `sig_method`, Type: `string`},
				},
			},
			{
				Name:        "登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。更多使用方法详见 小程序登录。",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/login/auth.code2Session.html",
				FuncName:    "Code2Session",
				GetParams: []Param{
					{Name: `appid`, Type: `string`},
					{Name: `secret`, Type: `string`},
					{Name: `js_code`, Type: `string`},
					{Name: `grant_type`, Type: `string`},
				},
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
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/sec-check/security.imgSecCheck.html",
				FuncName:    "ImgSecCheck",
			},
			{
				Name:        "异步校验图片/音频是否含有违法违规内容。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/media_check_async?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html",
				FuncName:    "MediaCheckAsync",
			},
			{
				Name:        "检查一段文本是否含有违法违规内容。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/msg_sec_check?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/sec-check/security.msgSecCheck.html",
				FuncName:    "MsgSecCheck",
			},
		},
	},
	{
		Name:    `开放数据`,
		Package: `data`,
		Apis: []Api{

			{
				Name:        "删除已经上报到微信的key-value数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/remove_user_storage?access_token=ACCESS_TOKEN&signature=SIGNATURE&openid=OPENID&sig_method=SIG_METHOD",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data/storage.removeUserStorage.html",
				FuncName:    "RemoveUserStorage",
				GetParams: []Param{
					{Name: `signature`, Type: `string`},
					{Name: `openid`, Type: `string`},
					{Name: `sig_method`, Type: `string`},
				},
			},
			{
				Name:        "写用户关系链互动数据存储",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/setuserinteractivedata?access_token=ACCESS_TOKEN&signature=SIGNATURE&openid=OPENID&sig_method=SIG_METHOD",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data/storage.setUserInteractiveData.html",
				FuncName:    "SetUserInteractiveData",
				GetParams: []Param{
					{Name: `signature`, Type: `string`},
					{Name: `openid`, Type: `string`},
					{Name: `sig_method`, Type: `string`},
				},
			},
			{
				Name:        "上报用户数据后台接口。小游戏可以通过本接口上报key-value数据到用户的CloudStorage。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/set_user_storage?access_token=ACCESS_TOKEN&signature=SIGNATURE&openid=OPENID&sig_method=SIG_METHOD",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data/storage.setUserStorage.html",
				FuncName:    "SetUserStorage",
				GetParams: []Param{
					{Name: `openid`, Type: `string`},
					{Name: `sig_method`, Type: `string`},
					{Name: `signature`, Type: `string`},
				},
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
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/updatable-message/updatableMessage.createActivityId.html",
				FuncName:    "CreateActivityId",
			},
			{
				Name:        "修改被分享的动态消息。详见动态消息。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/wxopen/updatablemsg/send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/updatable-message/updatableMessage.setUpdatableMsg.html",
				FuncName:    "SetUpdatableMsg",
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
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html",
				FuncName:    "CreateQRCode",
			},
			{
				Name:        "获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/qr-code/wxacode.get.html",
				FuncName:    "Get",
			},
			{
				Name:        "获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。 更多用法详见 获取二维码。",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html",
				FuncName:    "GetUnlimited",
			},
		},
	},
	{
		Name:    `帧同步`,
		Package: `frame`,
		Apis: []Api{

			{
				Name:        "第三方后台创建帧同步游戏房间",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/createwxagameroom?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.createGameRoom.html",
				FuncName:    "CreateGameRoom",
			},
			{
				Name:        "分片拉取对局游戏帧",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxa/getwxagameframe?access_token=XXX&access_info=YYY&begin_frame_id=ZZZ&end_frame_id=TTT",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.getGameFrame.html",
				FuncName:    "GetGameFrame",
				GetParams: []Param{
					{Name: `end_frame_id`, Type: `string`},
					{Name: `access_info`, Type: `string`},
					{Name: `begin_frame_id`, Type: `string`},
				},
			},
			{
				Name:        "获取对局玩家位次信息",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxa/getwxagameidentityinfo?access_token=XXX&access_info=YYY",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.getGameIdentityInfo.html",
				FuncName:    "GetGameIdentityInfo",
				GetParams: []Param{
					{Name: `access_info`, Type: `string`},
				},
			},
			{
				Name:        "获取指定房间信息",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxa/getwxagameroominfo?access_token=XXX&access_info=YYY",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.getGameRoomInfo.html",
				FuncName:    "GetGameRoomInfo",
				GetParams: []Param{
					{Name: `access_info`, Type: `string`},
				},
			},
		},
	},
	{
		Name:    `订阅消息`,
		Package: `subscribe`,
		Apis: []Api{

			{
				Name:        "发送订阅消息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html",
				FuncName:    "Send",
			},
		},
	},
}
