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

// HTTPAuthTypeDMethod URL鉴权方式D
// https://domainName/path?sign=hash&t=timestamp
// Hash算法：HASH(path@timestamp@secret)
// 回源URL：https://domainName/path
type HTTPAuthTypeDMethod struct {
	HTTPAuthBaseMethod

	SignParamName      string `json:"signParamName"`
	TimestampParamName string `json:"timestampParamName"`
	Secret             string `json:"secret"`
	Life               int    `json:"life"`
	Algorithm          string `json:"algorithm"` // Hash算法，暂时不开放 TODO
}

func NewHTTPAuthTypeDMethod() *HTTPAuthTypeDMethod {
	return &HTTPAuthTypeDMethod{
		SignParamName:      "sign",
		TimestampParamName: "t",
		Secret:             rands.HexString(16),
		Life:               30,
	}
}

// Init 初始化
func (this *HTTPAuthTypeDMethod) Init(params map[string]any) error {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}
	return json.Unmarshal(paramsJSON, this)
}

// Filter 过滤
func (this *HTTPAuthTypeDMethod) Filter(req *http.Request, subReqFunc func(subReq *http.Request) (status int, err error), formatter func(string) string) (ok bool, newURI string, uriChanged bool, err error) {
	var signParamName = this.SignParamName
	if len(signParamName) == 0 {
		signParamName = "sign"
	}

	var timestampParamName = this.TimestampParamName
	if len(timestampParamName) == 0 {
		timestampParamName = "t"
	}

	var life = this.Life
	if life <= 0 {
		life = 30
	}

	var path = req.URL.Path
	var query = req.URL.Query()
	var hash = query.Get(signParamName)
	if len(hash) == 0 {
		return false, "", false, nil
	}

	var timestamp = query.Get(timestampParamName)
	if !this.matchTimestamp(timestamp) {
		return false, "", false, nil
	}

	var timestampInt64 = types.Int64(timestamp)
	if timestampInt64 <= time.Now().Unix()-int64(life) {
		return false, "", false, nil
	}

	var realHash = fmt.Sprintf("%x", md5.Sum([]byte(path+"@"+timestamp+"@"+this.Secret)))
	if realHash != hash {
		return false, "", false, nil
	}

	// 去除参数
	var uri = req.URL.RequestURI()
	var qIndex = strings.Index(uri, "?")
	if qIndex >= 0 {
		path = uri[:qIndex]
		var args = this.removeQueryArgs(uri[qIndex+1:], []string{this.SignParamName, this.TimestampParamName})
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
