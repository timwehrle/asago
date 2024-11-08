package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ConfigCmd)
	RootCmd.AddCommand(UserCmd)
	RootCmd.AddCommand(TasksCmd)
}

var RootCmd = &cobra.Command{
	Use:   "asago",
	Short: "asago is a CLI tool for managing your Asana.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
