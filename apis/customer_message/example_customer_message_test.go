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

package customer_message_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/miniprogram"
	"github.com/fastwego/miniprogram/apis/customer_message"
)

func ExampleGetTempMedia() {
	var ctx *miniprogram.Miniprogram

	params := url.Values{}
	resp, err := customer_message.GetTempMedia(ctx, params)

	fmt.Println(resp, err)
}

func ExampleSend() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := customer_message.Send(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetTyping() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := customer_message.SetTyping(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUploadTempMedia() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	params := url.Values{}
	resp, err := customer_message.UploadTempMedia(ctx, payload, params)

	fmt.Println(resp, err)
}
