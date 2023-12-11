package main

import (
	"config-management-go/config"
	"fmt"

	_ "github.com/gin-gonic/gin"
)

func main() {
	// 读取配置文件
	config, err := config.ReadConfig("./config/local.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印 gorm 配置
	fmt.Println(config.Gorm.Driver)
	fmt.Println(config.Gorm.Dsn)
}
