package config

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/json"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
)

// LoadConfigFile 载入配置文件.
func LoadConfigFile(path string) (config.Config, error) {
	cfg, _ := config.NewConfig()
	if err := cfg.Load(file.NewSource(file.WithPath(path), source.WithEncoder(json.NewEncoder()))); err != nil {
		return nil, err
	}
	return cfg, nil
}

// MicroSrvConf MicroService基础配置.
type MicroSrvConf struct {
	Basic struct {
		Name             string `json:"name"`     // 服务名;
		Version          string `json:"version"`  // 服务版本;
		Address          string `json:"address"`  // 服务侦听地址;
		Registry         string `json:"register"` // 服务注册地址;
		RegisterTTL      int64  `json:"ttl"`
		RegisterInterval int64  `json:"interval"`
	} `json:"basic"`
}

// WithMicroSrvConf 使用MicroService基础配置.
func WithMicroSrvConf(cfg config.Config) (*MicroSrvConf, error) {
	var node = &MicroSrvConf{}
	if err := cfg.Scan(node); err != nil {
		return nil, err
	}
	return node, nil
}
