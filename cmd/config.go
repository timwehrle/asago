package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

const (
	service = "asago"
	username = "user"
)

var ConfigCmd = &cobra.Command{
	Use: "config",
	Short: "Manage the Asana configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := keyring.Get(service, username)
		if err != nil {
			fmt.Println("Enter your Asana Personal Access Token (PAT):")
			fmt.Scanln(&token)

			err = keyring.Set(service, username, token)
			if err != nil {
				log.Fatalf("Failed to store token: %v", err)
			}
		}
	},
}