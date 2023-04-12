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

// HTTPAuthTypeBMethod URL鉴权方式B
// https://domainName/timestamp/hash/path
// Hash算法：HASH(path@timestamp@secret)
// 回源URL：https://domainName/path
type HTTPAuthTypeBMethod struct {
	HTTPAuthBaseMethod

	Secret    string `json:"secret"`
	Life      int    `json:"life"`
	Algorithm string `json:"algorithm"` // Hash算法，暂时不开放 TODO
}

func NewHTTPAuthTypeBMethod() *HTTPAuthTypeBMethod {
	return &HTTPAuthTypeBMethod{
		Secret: rands.HexString(16),
		Life:   30,
	}
}

// Init 初始化
func (this *HTTPAuthTypeBMethod) Init(params map[string]any) error {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}
	return json.Unmarshal(paramsJSON, this)
}

// Filter 过滤
func (this *HTTPAuthTypeBMethod) Filter(req *http.Request, subReqFunc func(subReq *http.Request) (status int, err error), formatter func(string) string) (ok bool, newURI string, uriChanged bool, err error) {
	var life = this.Life
	if life <= 0 {
		life = 30
	}

	var path = req.URL.Path
	var pieces = strings.SplitN(path, "/", 4)
	if len(pieces) != 4 {
		return false, "", false, nil
	}

	// pieces[0] = "/"
	var timestamp = pieces[1]
	var hash = pieces[2]
	path = "/" + pieces[3]

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

	// 修改路径
	var query = req.URL.RawQuery
	if len(query) > 0 {
		newURI = path + "?" + query
	} else {
		newURI = path
	}

	return true, newURI, true, nil
}
