package systemconfigs

import "github.com/1uLang/EdgeCommon/pkg/serverconfigs/shared"

// 默认日志配置
func DefaultLogConfig() *LogConfig {
	return &LogConfig{
		CanDelete: false,
		CanClean:  false,
		Capacity: &shared.SizeCapacity{
			Count: 1,
			Unit:  shared.SizeCapacityUnitGB,
		},
		Days:      30,
		CanChange: true,
	}
}

// 操作日志相关配置
type LogConfig struct {
	CanDelete bool                 `json:"canDelete"` // 是否可删除
	CanClean  bool                 `json:"canClean"`  // 是否可清理
	Capacity  *shared.SizeCapacity `json:"capacity"`  // 容量
	Days      int                  `json:"days"`      // 自动保存天数
	CanChange bool                 `json:"canChange"` // 是否允许再次修改配置
}
