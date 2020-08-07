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

// Package security 内容安全
package security

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/fastwego/miniprogram"
)

const (
	apiImgSecCheck     = "/wxa/img_sec_check"
	apiMediaCheckAsync = "/wxa/media_check_async"
	apiMsgSecCheck     = "/wxa/msg_sec_check"
)

/*
校验一张图片是否含有违法违规内容。



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/sec-check/security.imgSecCheck.html

POST(@media) https://api.weixin.qq.com/wxa/img_sec_check?access_token=ACCESS_TOKEN
*/
func ImgSecCheck(ctx *miniprogram.Miniprogram, media string) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

	}()
	return ctx.Client.HTTPPost(apiImgSecCheck, r, m.FormDataContentType())
}

/*
异步校验图片/音频是否含有违法违规内容。



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html

POST https://api.weixin.qq.com/wxa/media_check_async?access_token=ACCESS_TOKEN
*/
func MediaCheckAsync(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMediaCheckAsync, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
检查一段文本是否含有违法违规内容。



See: https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/sec-check/security.msgSecCheck.html

POST https://api.weixin.qq.com/wxa/msg_sec_check?access_token=ACCESS_TOKEN
*/
func MsgSecCheck(ctx *miniprogram.Miniprogram, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMsgSecCheck, bytes.NewReader(payload), "application/json;charset=utf-8")
}
