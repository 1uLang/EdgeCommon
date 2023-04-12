// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package userconfigs

type UserRegisterConfig struct {
	IsOn            bool `yaml:"isOn" json:"isOn"`                       // 是否启用用户注册
	ComplexPassword bool `yaml:"complexPassword" json:"complexPassword"` // 必须使用复杂密码

	// CDN
	CDNIsOn             bool     `json:"cdnIsOn"`                                        // 是否开启CDN服务
	ClusterId           int64    `yaml:"clusterId" json:"clusterId"`                     // 用户创建服务集群
	Features            []string `yaml:"features" json:"features"`                       // 默认启用的功能
	RequireVerification bool     `yaml:"requireVerification" json:"requireVerification"` // 是否需要审核
	RequireIdentity     bool     `yaml:"requireIdentity" json:"requireIdentity"`         // 是否需要实名认证

	// 开通DNS
	NSIsOn bool `json:"nsIsOn"` // 是否开启智能DNS服务
}

func DefaultUserRegisterConfig() *UserRegisterConfig {
	return &UserRegisterConfig{
		IsOn:            false,
		ComplexPassword: true,
		CDNIsOn:         true,
		NSIsOn:          false,
		Features: []string{
			UserFeatureCodeServerAccessLog,
			UserFeatureCodeServerViewAccessLog,
			UserFeatureCodeServerWAF,
			UserFeatureCodePlan,
		},
		RequireVerification: false,
	}
}