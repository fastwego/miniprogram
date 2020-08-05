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

/*
微信小程序开发 SDK

See: https://developers.weixin.qq.com/miniprogram/dev/api-backend/
*/
package miniprogram

import (
	"log"
	"os"

	"github.com/faabiosr/cachego"
	"github.com/faabiosr/cachego/file"
)

// GetAccessTokenFunc 获取 access_token 方法接口
type GetAccessTokenFunc func(ctx *Miniprogram) (accessToken string, err error)

/*
Miniprogram 实例
*/
type Miniprogram struct {
	Config      MiniprogramConfig
	AccessToken AccessToken
	Client      Client
	Logger      *log.Logger
}

/*
AccessToken 管理器 处理缓存 和 刷新 逻辑
*/
type AccessToken struct {
	Cache                 cachego.Cache
	GetAccessTokenHandler GetAccessTokenFunc
}

/*
小程序配置
*/
type MiniprogramConfig struct {
	Appid  string
	Secret string
}

/*
创建小程序实例
*/
func New(config MiniprogramConfig) (miniprogram *Miniprogram) {
	instance := Miniprogram{
		Config: config,
		AccessToken: AccessToken{
			Cache:                 file.New(os.TempDir()),
			GetAccessTokenHandler: GetAccessToken,
		},
	}

	instance.Client = Client{Ctx: &instance}

	instance.Logger = log.New(os.Stdout, "[miniprogram] ", log.LstdFlags)

	return &instance
}

/*
SetAccessTokenCacheDriver 设置 AccessToken 缓存器 默认为文件缓存：目录 os.TempDir()

驱动接口类型 为 cachego.Cache
*/
func (miniprogram *Miniprogram) SetAccessTokenCacheDriver(driver cachego.Cache) {
	miniprogram.AccessToken.Cache = driver
}

/*
SetGetAccessTokenHandler 设置 AccessToken 获取方法。默认 从本地缓存获取（过期从微信接口刷新）

如果有多实例服务，可以设置为 Redis 或 RPC 等中控服务器 获取 就可以避免 AccessToken 刷新冲突
*/
func (miniprogram *Miniprogram) SetGetAccessTokenHandler(f GetAccessTokenFunc) {
	miniprogram.AccessToken.GetAccessTokenHandler = f
}

/*
SetLogger 日志记录 默认输出到 os.Stdout

可以新建 logger 输出到指定文件

如果不想开启日志，可以输出到 /dev/null log.SetOutput(ioutil.Discard)
*/
func (miniprogram *Miniprogram) SetLogger(logger *log.Logger) {
	miniprogram.Logger = logger
}
