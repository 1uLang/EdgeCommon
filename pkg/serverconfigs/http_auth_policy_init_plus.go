// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build plus

package serverconfigs

import (
	"errors"
)

// Init 初始化
func (this *HTTPAuthPolicy) Init() error {
	switch this.Type {
	case HTTPAuthTypeBasicAuth:
		this.method = NewHTTPAuthBasicMethod()
	case HTTPAuthTypeSubRequest:
		this.method = NewHTTPAuthSubRequestMethod()
	case HTTPAuthTypeTypeA:
		this.method = NewHTTPAuthTypeAMethod()
	case HTTPAuthTypeTypeB:
		this.method = NewHTTPAuthTypeBMethod()
	case HTTPAuthTypeTypeC:
		this.method = NewHTTPAuthTypeCMethod()
	case HTTPAuthTypeTypeD:
		this.method = NewHTTPAuthTypeDMethod()
	}

	if this.method == nil {
		return errors.New("unknown auth method '" + this.Type + "'")
	}
	err := this.method.Init(this.Params)
	if err != nil {
		return err
	}

	return nil
}
