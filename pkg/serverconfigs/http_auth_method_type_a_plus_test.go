// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build plus
// +build plus

package serverconfigs_test

import (
	"crypto/md5"
	"fmt"
	"github.com/1uLang/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/assert"
	"github.com/iwind/TeaGo/rands"
	"github.com/iwind/TeaGo/types"
	"net/http"
	"testing"
	"time"
)

func TestHTTPAuthTypeAMethod_Filter(t *testing.T) {
	var a = assert.NewAssertion(t)

	var path = "/images/test.jpg"
	var secret = "123456"
	var timestamp = types.String(time.Now().Unix())
	var randomString = rands.HexString(32)
	var hash = fmt.Sprintf("%x", md5.Sum([]byte(path+"@"+timestamp+"@"+randomString+"@"+secret)))

	req, err := http.NewRequest(http.MethodGet, "https://example.com/images/test.jpg?sign="+timestamp+"-"+randomString+"-"+hash+"&v=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	var method = serverconfigs.NewHTTPAuthTypeAMethod()
	err = method.Init(map[string]any{
		"secret": secret,
		"life":   10,
	})
	if err != nil {
		t.Fatal(err)
	}
	ok, uri, uriChanged, err := method.Filter(req, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	a.IsTrue(ok)
	if ok {
		t.Log(req.URL)
		t.Log("newURI:", uri, uriChanged)
	}
}

func TestHTTPAuthTypeAMethod_URL(t *testing.T) {
	var path = "/images/test.jpg"
	var secret = "123456"
	var timestamp = types.String(time.Now().Unix())
	var randomString = rands.HexString(16)
	t.Log("timestamp:", timestamp)
	t.Log("random:", randomString)
	var hash = fmt.Sprintf("%x", md5.Sum([]byte(path+"@"+timestamp+"@"+randomString+"@"+secret)))

	t.Log("https://example.com" + path + "?sign=" + timestamp + "-" + randomString + "-" + hash)
}

func TestHTTPAuthTypeAMethod_URL2(t *testing.T) {
	var path = "/get"
	var secret = "123456"
	var timestamp = types.String(time.Now().Unix())
	var randomString = rands.HexString(16)
	var hash = fmt.Sprintf("%x", md5.Sum([]byte(path+"@"+timestamp+"@"+randomString+"@"+secret)))

	t.Log("https://example.com/get?sign=" + timestamp + "-" + randomString + "-" + hash)
}
