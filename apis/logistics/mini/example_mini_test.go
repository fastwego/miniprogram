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
	"github.com/fastwego/miniprogram/apis/logistics/mini"
)

func ExampleBatchGetOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.BatchGetOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.AddOrder(ctx, payload)

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

func ExampleGetAllAccount() {
	var ctx *miniprogram.Miniprogram

	resp, err := mini.GetAllAccount(ctx)

	fmt.Println(resp, err)
}

func ExampleGetAllDelivery() {
	var ctx *miniprogram.Miniprogram

	resp, err := mini.GetAllDelivery(ctx)

	fmt.Println(resp, err)
}

func ExampleGetOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.GetOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetPath() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.GetPath(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetPrinter() {
	var ctx *miniprogram.Miniprogram

	resp, err := mini.GetPrinter(ctx)

	fmt.Println(resp, err)
}

func ExampleGetQuota() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.GetQuota(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTestUpdateOrder() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.TestUpdateOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdatePrinter() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := mini.UpdatePrinter(ctx, payload)

	fmt.Println(resp, err)
}
