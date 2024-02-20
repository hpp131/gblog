package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gblog",
	Short: "gblog is root command",
	Long: "gblog is root command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute(){
	rootCmd.AddCommand(initCmd, startCmd)
	if err := rootCmd.Execute();err != nil{
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}