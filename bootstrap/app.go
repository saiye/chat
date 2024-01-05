package bootstrap

import "chat/providers"

type App struct {
	container *providers.Container
	providers []providers.Provider
}

func NewApp() *App {
	container := providers.NewContainer()
	return &App{
		container: container,
		providers: []providers.Provider{
			providers.NewMysqlProvider(container),
		},
	}
}

// BootstrapAllProviders 初始化所有服务提供者
func (receiver *App) BootstrapAllProviders() {
	for _, provider := range receiver.providers {
		provider.Bootstrap()
	}
}

// RunAllProviders 启动所有服务提供者
func (receiver *App) RunAllProviders() {
	for _, provider := range receiver.providers {
		provider.Run()
	}
}

// ExitAllProviders 退出所有服务提供者
func (receiver *App) ExitAllProviders() {
	for _, provider := range receiver.providers {
		provider.Exit()
	}
}
