// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build plus

package serverconfigs

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/iwind/TeaGo/rands"
	"github.com/iwind/TeaGo/types"
	"net/http"
	"strings"
	"time"
)

// HTTPAuthTypeAMethod URL鉴权方式A
// https://domainName/path?sign=timestamp-randomString-hash
// Hash算法：HASH(path@timestamp@randomString@secret)
// 回源URL（仍然保留sign参数）：https://domainName/path?sign=...
type HTTPAuthTypeAMethod struct {
	HTTPAuthBaseMethod

	SignParamName string `json:"signParamName"`
	Secret        string `json:"secret"`
	Life          int    `json:"life"`
	Algorithm     string `json:"algorithm"` // Hash算法，暂时不开放 TODO
}

func NewHTTPAuthTypeAMethod() *HTTPAuthTypeAMethod {
	return &HTTPAuthTypeAMethod{
		SignParamName: "sign",
		Secret:        rands.HexString(16),
		Life:          30,
	}
}

// Init 初始化
func (this *HTTPAuthTypeAMethod) Init(params map[string]any) error {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}
	return json.Unmarshal(paramsJSON, this)
}

// Filter 过滤
func (this *HTTPAuthTypeAMethod) Filter(req *http.Request, subReqFunc func(subReq *http.Request) (status int, err error), formatter func(string) string) (ok bool, newURI string, uriChanged bool, err error) {
	var signParamName = this.SignParamName
	if len(signParamName) == 0 {
		signParamName = "sign"
	}

	var life = this.Life
	if life <= 0 {
		life = 30
	}

	var path = req.URL.Path
	var sign = req.URL.Query().Get(signParamName)
	if len(sign) == 0 {
		return false, "", false, nil
	}

	var pieces = strings.Split(sign, "-")
	if len(pieces) != 3 {
		return false, "", false, nil
	}

	var timestamp = pieces[0]
	var randomString = pieces[1]
	var hash = pieces[2]

	if !this.matchTimestamp(timestamp) {
		return false, "", false, nil
	}

	var timestampInt64 = types.Int64(timestamp)
	if timestampInt64 <= time.Now().Unix()-int64(life) {
		return false, "", false, nil
	}

	var realHash = fmt.Sprintf("%x", md5.Sum([]byte(path+"@"+timestamp+"@"+randomString+"@"+this.Secret)))
	if realHash != hash {
		return false, "", false, nil
	}

	// 去除参数
	var uri = req.URL.RequestURI()
	var qIndex = strings.Index(uri, "?")
	if qIndex >= 0 {
		path = uri[:qIndex]
		var args = this.removeQueryArgs(uri[qIndex+1:], []string{this.SignParamName})
		if len(args) > 0 {
			newURI = path + "?" + args
		} else {
			newURI = path
		}
	} else {
		newURI = uri
	}

	return true, newURI, true, nil
}
