// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vhost

import (
	"bytes"
	"io"
	"net/http"
	"os"

	logpkg "github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/version"
)

var NotFoundPagePath = ""

const (
	NotFound = `<!DOCTYPE html>
<html>
<head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8"><title>404 找不到此页面/(ㄒoㄒ)/~~</title><meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0"><meta http-equiv="X-UA-Compatible" content="ie=edge"><style>*{padding:0;margin:0}a{text-decoration:none}.notfoud-container .img-404{height:155px;background:url(https://me.heimaoba.cn/frp_404.png) center center no-repeat;-webkit-background-size:150px auto;margin-top:40px;margin-bottom:20px}.notfoud-container .notfound-p{line-height:22px;font-size:17px;padding-bottom:15px;border-bottom:1px solid #f6f6f6;text-align:center;color:#262b31}.notfoud-container .notfound-reason{color:#9ca4ac;font-size:13px;line-height:13px;text-align:left;width:210px;margin:0 auto}.notfoud-container .notfound-reason p{margin-top:13px}.notfoud-container .notfound-reason ul li{margin-top:10px;margin-left:36px}.notfoud-container .notfound-btn-container{margin:40px auto 0;text-align:center}.notfoud-container .notfound-btn-container .notfound-btn{display:inline-block;border:1px solid #ebedef;background-color:#239bf0;color:#fff;font-size:15px;border-radius:5px;text-align:center;padding:10px;line-height:16px;white-space:nowrap}</style></head><body>

    <div class="notfoud-container">
        <div class="img-404">
        </div>
        <p class="notfound-p">哎呀迷路了...</p>
        <div class="notfound-reason">
            <p>可能的原因：</p>
            <ul>
                <li>frpc没有启动成功</li>
                <li>域名不正确</li>
                <li>内网端不存在该页面</li>
                <li>我们的服务器被外星人劫持了</li>
            </ul>
        </div>
        
        <div class="notfound-reason">
            <p align="center"><a href="https://app.heimaoba.cn/">黑猫之家</a>
            <p align="center"><a href="https://cloud.tencent.com/act/cps/redirect?redirect=2446&cps_key=8700306c81b92541829f34d1973f51bb&from=console">【腾讯云】云服务器、云数据库、COS、CDN、短信等云产品特惠热卖中</a>
            <p align="center"><a href="https://www.bt.cn/?invite_code=M19za2l1eW8=">宝塔服务器面板，一键全能部署及管理，送你10850元礼包，点我领取</a>

        </div>

    </div></body>
</html>
`
)

func getNotFoundPageContent() []byte {
	var (
		buf []byte
		err error
	)
	if NotFoundPagePath != "" {
		buf, err = os.ReadFile(NotFoundPagePath)
		if err != nil {
			logpkg.Warn("read custom 404 page error: %v", err)
			buf = []byte(NotFound)
		}
	} else {
		buf = []byte(NotFound)
	}
	return buf
}

func NotFoundResponse() *http.Response {
	header := make(http.Header)
	header.Set("server", "frp/"+version.Full())
	header.Set("Content-Type", "text/html")

	content := getNotFoundPageContent()
	res := &http.Response{
		Status:        "Not Found",
		StatusCode:    404,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        header,
		Body:          io.NopCloser(bytes.NewReader(content)),
		ContentLength: int64(len(content)),
	}
	return res
}
