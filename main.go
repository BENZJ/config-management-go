package main

import (
	"config-management-go/config"
	"config-management-go/injector"
)

func main() {
	// 读取配置文件
	config.MustLoad("./config/local.yaml")

	app, cleanup, err := injector.BuildApiInjector()
	if err != nil {
		panic(err)
	}
	defer cleanup()
	app.Engine.Run()
	// fmt.Print(app)
}
