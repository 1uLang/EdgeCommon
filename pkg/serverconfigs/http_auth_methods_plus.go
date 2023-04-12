// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build plus

package serverconfigs

type HTTPAuthType = string

const (
	HTTPAuthTypeBasicAuth  HTTPAuthType = "basicAuth"  // BasicAuth
	HTTPAuthTypeSubRequest HTTPAuthType = "subRequest" // 子请求

	HTTPAuthTypeTypeA HTTPAuthType = "typeA"
	HTTPAuthTypeTypeB HTTPAuthType = "typeB"
	HTTPAuthTypeTypeC HTTPAuthType = "typeC"
	HTTPAuthTypeTypeD HTTPAuthType = "typeD"
)

type HTTPAuthTypeDefinition struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	IsPlus      bool   `json:"isPlus"`
}

func FindAllHTTPAuthTypes() []*HTTPAuthTypeDefinition {
	return []*HTTPAuthTypeDefinition{
		{
			Name:        "URL鉴权A",
			Code:        HTTPAuthTypeTypeA,
			Description: "示例URL：https://example.com/images/test.jpg?sign=1661753445-8453b620246d423d-8b7221a273fa96d467052b421d1a9f37<br/>",
			IsPlus:      true,
		},
		{
			Name:        "URL鉴权B",
			Code:        HTTPAuthTypeTypeB,
			Description: "示例URL：https://example.com/1661753540/af5df08679b86f9a51436983eb86f820/images/test.jpg<br/>",
			IsPlus:      true,
		},
		{
			Name:        "URL鉴权C",
			Code:        HTTPAuthTypeTypeC,
			Description: "示例URL：https://example.com/d8cd2b671fae56342fdd82e3df88c9a3/1661753633/images/test.jpg<br/>",
			IsPlus:      true,
		},
		{
			Name:        "URL鉴权D",
			Code:        HTTPAuthTypeTypeD,
			Description: "示例URL：https://example.com/images/test.jpg?sign=f66af42f87cf63a64f4b86ec11c7797a&t=1661753717<br/>",
			IsPlus:      true,
		},
		{
			Name:        "基本认证",
			Code:        HTTPAuthTypeBasicAuth,
			Description: "BasicAuth，最简单的HTTP请求认证方式，通过传递<span class=\"ui label tiny basic text\">Authorization: Basic xxx</span> Header认证。",
		},
		{
			Name:        "子请求",
			Code:        HTTPAuthTypeSubRequest,
			Description: "通过自定义的URL子请求来认证请求。",
		},
	}
}
