package main

import (
	"config-management-go/config"
	"fmt"

	_ "github.com/gin-gonic/gin"
)

func main() {
	// 读取配置文件
	config.MustLoad("./config/local.yaml")

	// 打印 gorm 配置
	fmt.Println(config.C.Gorm.Driver)
	fmt.Println(config.C.Gorm.Dsn)
}
