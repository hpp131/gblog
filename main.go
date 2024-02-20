package main

import (
	// "github.com/gin-gonic/gin"
	_ "github.com/hpp131/gblog/apps" // 执行apps/init.go文件，运行所有ioc对象的Registry方法
	"github.com/hpp131/gblog/cmd"
	// "github.com/hpp131/gblog/ioc"
)

func main() {
	cmd.Execute()
}
