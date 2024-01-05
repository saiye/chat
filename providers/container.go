package providers

import (
	"fmt"
	"reflect"
)

// Container 结构体表示服务容器
type Container struct {
	services map[string]interface{}
}

// NewContainer 创建一个新的服务容器实例
func NewContainer() *Container {
	return &Container{
		services: make(map[string]interface{}),
	}
}

// Bind 将一个服务绑定到容器中
func (c *Container) Bind(name string, service interface{}) {
	c.services[name] = service
}

// Resolve 获取容器中的服务
func (c *Container) Resolve(name string) (interface{}, error) {
	service, ok := c.services[name]
	if !ok {
		return nil, fmt.Errorf("service %v not found in the container", name)
	}

	fnType := reflect.TypeOf(service)
	if fnType.Kind() == reflect.Func {
		// 如果是函数，尝试调用并返回结果
		result := reflect.ValueOf(service).Call(nil)

		// 检查是否有错误返回值
		if len(result) > 1 && !result[1].IsNil() {
			return nil, result[1].Interface().(error)
		}

		return result[0].Interface(), nil
	}

	return service, nil
}
