package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hpp131/gblog/ioc"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
	Short: "gblog start",
	Run: func(cmd *cobra.Command, args []string) {
	// 直接把原来main.go中的内容拿过来
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
	},
}