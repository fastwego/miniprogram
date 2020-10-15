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

package qrcode_test

import (
	"fmt"

	"github.com/fastwego/miniprogram"
	"github.com/fastwego/miniprogram/wxopen/apis/qrcode"
)

func ExampleQRCodeJumpGet() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := qrcode.QRCodeJumpGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQRCodeJumpDownload() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := qrcode.QRCodeJumpDownload(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQRCodeJumpAdd() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := qrcode.QRCodeJumpAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQRCodeJumpPublish() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := qrcode.QRCodeJumpPublish(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQRCodeJumpDelete() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := qrcode.QRCodeJumpDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleShorturl() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := qrcode.Shorturl(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetWxaCodeUnLimit() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := qrcode.GetWxaCodeUnLimit(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetWxaCode() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := qrcode.GetWxaCode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCreateWxaQRCode() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := qrcode.CreateWxaQRCode(ctx, payload)

	fmt.Println(resp, err)
}
