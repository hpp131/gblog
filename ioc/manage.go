package ioc

import "github.com/gin-gonic/gin"

// 多个namespace的ioc，使用map嵌套来实现

var nc = &NamespacedContainer{
	namespace: map[string]*Container{
		// 用于api层对象的注册， 如tokens.TokenAPIHandler对象
		"api": NewContainer(),
		// 用于controller层对象的注册， 如TokenServiceImpl对象
		"controller": NewContainer(),
		// 默认对象的注册
		"default": NewContainer(),
		// 用于配置对象的注册
		"config": NewContainer(),
	},
}

type NamespacedContainer struct {
	// 多个namespace
	namespace map[string]*Container
}

// 返回名为"api"的名称空间下的ioc container
func  API() *Container {
	return nc.namespace["api"]
}

// 返回名为"controller"的名称空间下的ioc container
func  Controller() *Container {
	return nc.namespace["controller"]
}

// func  Default() *Container{
// 	return nc.namespace["default"]
// }

// func  Config() *Container{
// 	return nc.namespace["config"]
// }

func  Init() error {
	// 循环执行ioc容器内已经注册的对象的Init方法
	for _, container := range nc.namespace{
		for _, obj := range container.storage{
			if err := obj.Init(); err != nil{
				return err
			}
		}
	}
	return nil
}

func Destroy() error {
	// 循环执行ioc容器内已经注册的对象的Destroy方法
	for _, container := range nc.namespace{
		for _, obj := range container.storage{
			if err := obj.Destroy(); err != nil{
				return err
			}
		}
	}
	return nil
}

func GinAPIRegistry(rr gin.IRouter) {
	// 循环执行ioc容器内已经注册的API对象的Registry方法
	for _, container := range nc.namespace{
		for _, obj := range container.storage{
			if api, ok := obj.(GinAPI); ok {
				api.Registry(rr)
			}
		}
	}
}


// 为APIHander提供根路由(rootrouter,简称rr)，如main.go中的TokenAPIHandler.Registry(rr)
type GinAPI interface{
	Objector
	Registry(gin.IRouter)
}
