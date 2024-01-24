package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hpp131/gblog/apps/tokens/api"
	"github.com/hpp131/gblog/apps/tokens/impl"
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