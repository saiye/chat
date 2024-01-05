package main

import "chat/bootstrap"

func main() {
	app := bootstrap.NewApp()
	// 销毁资源
	defer app.ExitAllProviders()
	// 引导程序
	app.BootstrapAllProviders()
	// 创建资源
	app.RunAllProviders()
}
