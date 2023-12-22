package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	C    = new(Config)
	once sync.Once
)

// 配置结构体
type Config struct {
	Gorm      Gorm      `yaml:"gorm"`
	WebConfig WebConfig `yaml:"webConfig"`
}

type Gorm struct {
	Driver string `yaml:"driver"`
	Dsn    string `yaml:"dsn"`
}
type WebConfig struct {
	Port string `yaml:"port"`
}

func MustLoad(filename string) {
	readConfig(filename)
}

// 读取配置
func readConfig(filename string) error {
	// 打开配置文件
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// 解析配置文件
	err = yaml.NewDecoder(f).Decode(C)
	if err != nil {
		return err
	}

	return nil
}
