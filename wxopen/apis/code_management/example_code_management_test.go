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

package code_management_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/miniprogram"
	"github.com/fastwego/miniprogram/wxopen/apis/code_management"
)

func ExampleCommit() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := code_management.Commit(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetPage() {
	var ctx *miniprogram.Miniprogram

	resp, err := code_management.GetPage(ctx)

	fmt.Println(resp, err)
}

func ExampleGetQrcode() {
	var ctx *miniprogram.Miniprogram

	params := url.Values{}
	resp, err := code_management.GetQrcode(ctx, params)

	fmt.Println(resp, err)
}

func ExampleSubmitAudit() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := code_management.SubmitAudit(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetAuditstatus() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := code_management.GetAuditstatus(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetLatestAuditstatus() {
	var ctx *miniprogram.Miniprogram

	resp, err := code_management.GetLatestAuditstatus(ctx)

	fmt.Println(resp, err)
}

func ExampleUndoCodeAudit() {
	var ctx *miniprogram.Miniprogram

	resp, err := code_management.UndoCodeAudit(ctx)

	fmt.Println(resp, err)
}

func ExampleRelease() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := code_management.Release(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleRevertCodeRelease() {
	var ctx *miniprogram.Miniprogram

	resp, err := code_management.RevertCodeRelease(ctx)

	fmt.Println(resp, err)
}

func ExampleGrayRelease() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := code_management.GrayRelease(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGrayReleasePlan() {
	var ctx *miniprogram.Miniprogram

	resp, err := code_management.GetGrayReleasePlan(ctx)

	fmt.Println(resp, err)
}

func ExampleRevertGrayRelease() {
	var ctx *miniprogram.Miniprogram

	resp, err := code_management.RevertGrayRelease(ctx)

	fmt.Println(resp, err)
}

func ExampleChangeVisitstatus() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := code_management.ChangeVisitstatus(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetWeAppSupportVersion() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := code_management.GetWeAppSupportVersion(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetWeAppSupportVersion() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := code_management.SetWeAppSupportVersion(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQueryQuota() {
	var ctx *miniprogram.Miniprogram

	resp, err := code_management.QueryQuota(ctx)

	fmt.Println(resp, err)
}

func ExampleSpeedupAudit() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := code_management.SpeedupAudit(ctx, payload)

	fmt.Println(resp, err)
}
