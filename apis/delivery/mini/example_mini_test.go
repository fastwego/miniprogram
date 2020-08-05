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

package mini_test

import (
	"fmt"

	"github.com/fastwego/miniprogram"
	"github.com/fastwego/miniprogram/apis/delivery/mini"
)

func ExampleAbnormalConfirm() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.AbnormalConfirm(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.AddOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddTip() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.AddTip(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBindAccount() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.BindAccount(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCancelOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.CancelOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetAllImmeDelivery() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.GetAllImmeDelivery(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetBindAccount() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.GetBindAccount(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.GetOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMockUpdateOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.MockUpdateOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleOpenDelivery() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.OpenDelivery(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePreAddOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.PreAddOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePreCancelOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.PreCancelOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleRealMockUpdateOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.RealMockUpdateOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.ReOrder(ctx, payload)

	fmt.Println(resp, err)
}
