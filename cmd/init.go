package cmd

import (
	"context"
	"fmt"

	"github.com/hpp131/gblog/apps/users"
	"github.com/hpp131/gblog/ioc"
	"github.com/spf13/cobra"
)

var (
	ctx      = context.Background()
	username string
	password string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "gblog init",
	Long:  "gblog init",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ioc.Init(); err != nil {
			panic(err)
		}
		// cmd.Help()
		// gblog init用来为系统创建Admin/管理员用户
		req := users.NewCreateUserRequest()
		// 可以使用cobra的flag功能在命令行中使用选项的方式来灵活定义username和password，避免硬编码
		req.Username = username
		req.Password = password
		req.Role = users.ADMIN

		user := ioc.Controller().Get(users.AppName).(users.Service)
		_, err := user.CreateUser(ctx, req)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Superadmin created successfully!")
		}
	},
}
