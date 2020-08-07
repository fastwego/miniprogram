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

package frame_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/miniprogram"
	"github.com/fastwego/miniprogram/minigame/apis/frame"
)

func ExampleCreateGameRoom() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := frame.CreateGameRoom(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGameFrame() {
	var ctx *miniprogram.Miniprogram

	params := url.Values{}
	resp, err := frame.GetGameFrame(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetGameIdentityInfo() {
	var ctx *miniprogram.Miniprogram

	params := url.Values{}
	resp, err := frame.GetGameIdentityInfo(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetGameRoomInfo() {
	var ctx *miniprogram.Miniprogram

	params := url.Values{}
	resp, err := frame.GetGameRoomInfo(ctx, params)

	fmt.Println(resp, err)
}
