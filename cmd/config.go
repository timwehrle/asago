package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/timwehrle/asago/api"
	"github.com/timwehrle/asago/api/endpoints"
	"github.com/timwehrle/asago/pkg/asago"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage the Asana configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := asago.GetToken()
		if err != nil {
			fmt.Println("Enter your Asana Personal Access Token (PAT):")
			fmt.Scanln(&token)

			err = asago.SetToken(token)
			if err != nil {
				log.Fatalf("Failed to store token: %v", err)
			}
			asago.AppConfig.DeleteToken = false

			client := api.NewClient(token)
			meEndpoint := endpoints.Me(client)

			user, err := meEndpoint.Get()
			if err != nil {
				log.Fatalf("Failed to get user info: %v", err)
			}

			huh.NewSelect[string]().
				Title("Choose a default workspace").
				Options(
					huh.NewOption(user.Workspaces[0].Name, user.Workspaces[0].Gid),
					huh.NewOption(user.Workspaces[1].Name, user.Workspaces[1].Gid),
				).
				Value(&asago.AppConfig.DefaultWorkspace).
				Run()

			if err := asago.SaveConfig(); err != nil {
				log.Fatalf("Failed to save config: %v", err)
			}
			fmt.Println("Default workspace set to", asago.AppConfig.DefaultWorkspace)
		} else {
			huh.NewConfirm().
				Title("Do you want to delete the stored token?").
				Affirmative("Yes").
				Negative("No").
				Value(&asago.AppConfig.DeleteToken).
				Run()

			if asago.AppConfig.DeleteToken {
				err := asago.DeleteToken()
				if err != nil {
					log.Fatalf("Failed to delete token: %v", err)
				}
				fmt.Println("Token deleted.")
			}
		}
	},
}
