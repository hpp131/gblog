package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/hpp131/gblog/apps"	// 执行apps/init.go文件，运行所有ioc对象的Registry方法
	"github.com/hpp131/gblog/ioc"
)

func main() {
	// 进行ioc容器内所有对象的初始化（即依赖解决）
	if err := ioc.Init();err != nil{
		panic(err)
	}
	engine := gin.Default()
	rootRouter := engine.Group("/gblogs/api/v1")
	ioc.GinAPIRegistry(rootRouter)
	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
