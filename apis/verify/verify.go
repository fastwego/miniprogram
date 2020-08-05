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

// Package verify 生物认证
package verify

import (
	"bytes"

	"github.com/fastwego/miniprogram"
)

const (
	apiVerifySignature = "/cgi-bin/soter/verify_signature"
)

/*
SOTER 生物认证秘钥签名验证



See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/soter/soter.verifySignature.html

POST https://api.weixin.qq.com/cgi-bin/soter/verify_signature?access_token=ACCESS_TOKEN
*/
func VerifySignature(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiVerifySignature, bytes.NewReader(payload), "application/json;charset=utf-8")
}
