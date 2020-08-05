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

package datacube_test

import (
	"fmt"

	"github.com/fastwego/miniprogram"
	"github.com/fastwego/miniprogram/apis/datacube"
)

func ExampleGetDailyRetain() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := datacube.GetDailyRetain(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMonthlyRetain() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := datacube.GetMonthlyRetain(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetWeeklyRetain() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := datacube.GetWeeklyRetain(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetDailySummary() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := datacube.GetDailySummary(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetDailyVisitTrend() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := datacube.GetDailyVisitTrend(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMonthlyVisitTrend() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := datacube.GetMonthlyVisitTrend(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetWeeklyVisitTrend() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := datacube.GetWeeklyVisitTrend(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserPortrait() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := datacube.GetUserPortrait(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetVisitDistribution() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := datacube.GetVisitDistribution(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetVisitPage() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := datacube.GetVisitPage(ctx, payload)

	fmt.Println(resp, err)
}
