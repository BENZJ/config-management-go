package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// 配置结构体
type Config struct {
	Gorm struct {
		Driver string `yaml:"driver"`
		Dsn    string `yaml:"dsn"`
	} `yaml:"gorm"`
	SqlLite struct {
		Path string `yaml:"path"`
	} `yaml:"sqlLite"`
}

// 读取配置
func ReadConfig(filename string) (*Config, error) {
	// 打开配置文件
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 解析配置文件
	config := &Config{}
	err = yaml.NewDecoder(f).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
