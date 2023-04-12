// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package dnsconfigs

import "github.com/1uLang/EdgeCommon/pkg/serverconfigs/shared"

type NSAnswerMode = string

const (
	NSAnswerModeRandom     NSAnswerMode = "random"
	NSAnswerModeRoundRobin NSAnswerMode = "roundRobin"
)

const NSAnswerDefaultSize = 5

type NSAnswerConfig struct {
	Mode    NSAnswerMode `yaml:"mode" json:"mode"`       // 记录回复模式
	MaxSize int16        `yaml:"maxSize" json:"maxSize"` // 记录回复最大数量
}

func (this *NSAnswerConfig) IsSame(config2 *NSAnswerConfig) bool {
	if config2 == nil {
		return false
	}

	return this.Mode == config2.Mode &&
		this.MaxSize == config2.MaxSize
}

func DefaultNSAnswerConfig() *NSAnswerConfig {
	return &NSAnswerConfig{
		Mode:    NSAnswerModeRandom,
		MaxSize: NSAnswerDefaultSize,
	}
}

func FindAllNSAnswerModes() []*shared.Definition {
	return []*shared.Definition{
		{
			Name:        "随机",
			Code:        NSAnswerModeRandom,
			Description: "有多个查询结果时，随机选取若干返回。",
		},
		{
			Name:        "轮询",
			Code:        NSAnswerModeRoundRobin,
			Description: "有多个查询结果时，按顺序每次返回一个。",
		},
	}
}

func IsValidNSAnswerMode(mode string) bool {
	for _, m := range FindAllNSAnswerModes() {
		if m.Code == mode {
			return true
		}
	}
	return false
}
