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

package subscribe_component_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/miniprogram"
	"github.com/fastwego/miniprogram/wxopen/apis/subscribe_component"
)

func ExampleGetShowWxaItem() {
	var ctx *miniprogram.Miniprogram

	resp, err := subscribe_component.GetShowWxaItem(ctx)

	fmt.Println(resp, err)
}

func ExampleGetWxaMplinkForShow() {
	var ctx *miniprogram.Miniprogram

	params := url.Values{}
	resp, err := subscribe_component.GetWxaMplinkForShow(ctx, params)

	fmt.Println(resp, err)
}

func ExampleUpdateShowWxaItem() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := subscribe_component.UpdateShowWxaItem(ctx, payload)

	fmt.Println(resp, err)
}
