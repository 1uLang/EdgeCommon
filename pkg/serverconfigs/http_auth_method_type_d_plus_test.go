// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build plus

package serverconfigs_test

import (
	"crypto/md5"
	"fmt"
	"github.com/1uLang/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/assert"
	"github.com/iwind/TeaGo/types"
	"net/http"
	"testing"
	"time"
)

func TestHTTPAuthTypeDMethod_Filter(t *testing.T) {
	var a = assert.NewAssertion(t)

	var path = "/images/test.jpg"
	var secret = "123456"
	var timestamp = types.String(time.Now().Unix())
	var hash = fmt.Sprintf("%x", md5.Sum([]byte(path+"@"+timestamp+"@"+secret)))

	req, err := http.NewRequest(http.MethodGet, "https://example.com/images/test.jpg?sign="+hash+"&t="+timestamp+"&v=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	var method = serverconfigs.NewHTTPAuthTypeDMethod()
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

func TestHTTPAuthTypeDMethod_URL(t *testing.T) {
	var path = "/images/test.jpg"
	var secret = "123456"
	var timestamp = types.String(time.Now().Unix())
	var hash = fmt.Sprintf("%x", md5.Sum([]byte(path+"@"+timestamp+"@"+secret)))
	t.Log("https://example.com/images/test.jpg?sign=" + hash + "&t=" + timestamp)
}
