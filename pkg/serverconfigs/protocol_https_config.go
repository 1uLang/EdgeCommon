package serverconfigs

import (
	"encoding/json"
	"github.com/1uLang/EdgeCommon/pkg/serverconfigs/sslconfigs"
)

func NewHTTPSProtocolConfigFromJSON(configJSON []byte) (*HTTPSProtocolConfig, error) {
	config := &HTTPSProtocolConfig{}
	if len(configJSON) > 0 {
		err := json.Unmarshal(configJSON, config)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

// HTTPS协议配置
type HTTPSProtocolConfig struct {
	BaseProtocol `yaml:",inline"`

	SSLPolicyRef *sslconfigs.SSLPolicyRef `yaml:"sslPolicyRef" json:"sslPolicyRef"`
	SSLPolicy    *sslconfigs.SSLPolicy    `yaml:"sslPolicy" json:"sslPolicy"`
}

// 初始化
func (this *HTTPSProtocolConfig) Init() error {
	err := this.InitBase()
	if err != nil {
		return err
	}

	if this.SSLPolicy != nil {
		err := this.SSLPolicy.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 转换为JSON
func (this *HTTPSProtocolConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
