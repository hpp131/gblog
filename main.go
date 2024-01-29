package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hpp131/gblog/apps/tokens/api"
	"github.com/hpp131/gblog/apps/tokens/impl"
	// 执行apps/init.go文件，运行所有ioc对象的Registry方法
	_ "github.com/hpp131/gblog/apps"
)



func main()  {
	tokenServiceImpl := impl.NewTokenServiceImpl()
	tokenAPIHandler := api.NewTokenAPIHandler(tokenServiceImpl)
	engine := gin.Default()
	rootRouter := engine.Group("/gblogs/api/v1")
	tokenAPIHandler.Registry(rootRouter)
	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}	
}