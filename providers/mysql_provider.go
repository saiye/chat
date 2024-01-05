package providers

import (
	"fmt"
)

type MysqlProvider struct {
	container *Container
}

func NewMysqlProvider(container *Container) *MysqlProvider {
	return &MysqlProvider{
		container: container,
	}
}
func (receiver *MysqlProvider) Bootstrap() {
	//receiver.container.Bind("mysql", func() {
	//	return service.MysqlCon()
	//})
}
func (receiver *MysqlProvider) Run() {
	fmt.Println("mysql Run")

}
func (receiver *MysqlProvider) Exit() {
	fmt.Println("mysql Exit")
}
