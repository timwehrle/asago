package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/timwehrle/asago/api"
	"github.com/timwehrle/asago/api/endpoints"
	"github.com/timwehrle/asago/pkg/asago"
)

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "Show the user information.",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := asago.GetToken()
		if err != nil {
			log.Fatalf("failed to get token: %v", err)
		}

		client := api.NewClient(token)
		meEndpoint := endpoints.Me(client)

		user, err := meEndpoint.Get()
		if err != nil {
			log.Fatalf("failed to get user info: %v", err)
		}

		fmt.Printf("Name: %s\nEmail: %s\n", user.Name, user.Email)
	},
}
