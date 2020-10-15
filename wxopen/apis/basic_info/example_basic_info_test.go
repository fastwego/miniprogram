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

package basic_info_test

import (
	"fmt"

	"github.com/fastwego/miniprogram"
	"github.com/fastwego/miniprogram/wxopen/apis/basic_info"
)

func ExampleGetAccountBasicInfo() {
	var ctx *miniprogram.Miniprogram

	resp, err := basic_info.GetAccountBasicInfo(ctx)

	fmt.Println(resp, err)
}

func ExampleModifyDomain() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := basic_info.ModifyDomain(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetWebviewDomain() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := basic_info.SetWebviewDomain(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetNickname() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := basic_info.SetNickname(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCheckWXVerifyNickname() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := basic_info.CheckWXVerifyNickname(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQueryNickname() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := basic_info.QueryNickname(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleModifyHeadImage() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := basic_info.ModifyHeadImage(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleModifySignature() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := basic_info.ModifySignature(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetWxaSearchStatus() {
	var ctx *miniprogram.Miniprogram

	resp, err := basic_info.GetWxaSearchStatus(ctx)

	fmt.Println(resp, err)
}

func ExampleChangeWxaSearchStatus() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := basic_info.ChangeWxaSearchStatus(ctx, payload)

	fmt.Println(resp, err)
}
