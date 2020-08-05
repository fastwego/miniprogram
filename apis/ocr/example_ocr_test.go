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

package ocr_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/miniprogram"
	"github.com/fastwego/miniprogram/apis/ocr"
)

func ExampleBankcard() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	params := url.Values{}
	resp, err := ocr.Bankcard(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleBusinessLicense() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	params := url.Values{}
	resp, err := ocr.BusinessLicense(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleDriverLicense() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	params := url.Values{}
	resp, err := ocr.DriverLicense(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleIdcard() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	params := url.Values{}
	resp, err := ocr.Idcard(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExamplePrintedText() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	params := url.Values{}
	resp, err := ocr.PrintedText(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleVehicleLicense() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	params := url.Values{}
	resp, err := ocr.VehicleLicense(ctx, payload, params)

	fmt.Println(resp, err)
}
