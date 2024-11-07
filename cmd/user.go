package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timwehrle/asago/pkg/asago"
)

var UserCmd = &cobra.Command{
	Use: "user",
	Short: "Show the user information.",
	Run: func(cmd *cobra.Command, args []string) {
		asago.FetchUserInfo()
	},
}