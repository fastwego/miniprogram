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

// Package member_admin 成员管理
package member_admin

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiBindTester   = "/wxa/bind_tester"
	apiUnbindTester = "/wxa/unbind_tester"
	apiMemberAuth   = "/wxa/memberauth"
)

/*
绑定微信用户为体验者

第三方平台在帮助旗下授权的小程序提交代码审核之前，可先让小程序运营者体验，体验之前需要将运营者的个人微信号添加到该小程序的体验者名单中。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Admin.html

POST https://api.weixin.qq.com/wxa/bind_tester?access_token=ACCESS_TOKEN
*/
func BindTester(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBindTester, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
解除绑定体验者

调用本接口可以将特定微信用户从小程序的体验者列表中解绑。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/unbind_tester.html

POST https://api.weixin.qq.com/wxa/unbind_tester?access_token=ACCESS_TOKEN
*/
func UnbindTester(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUnbindTester, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取体验者列表

通过本接口可以获取小程序所有已绑定的体验者列表。使用过程中如遇到问题，可在开放平台服务商专区发帖交流。

See: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/memberauth.html

POST https://api.weixin.qq.com/wxa/memberauth?access_token=ACCESS_TOKEN
*/
func MemberAuth(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMemberAuth, bytes.NewReader(payload), "application/json;charset=utf-8")
}
