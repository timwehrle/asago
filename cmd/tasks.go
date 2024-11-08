package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/timwehrle/asago/api"
	"github.com/timwehrle/asago/api/endpoints"
	"github.com/timwehrle/asago/pkg/asago"
)

var TasksCmd = &cobra.Command{
	Use: "tasks",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := asago.GetToken()
		if err != nil {
			log.Fatalf("failed to get token: %v", err)
		}	

		client := api.NewClient(token)
		tasksEndpoint := endpoints.Tasks(client)

		tasks, err := tasksEndpoint.ListTasks()
		if err != nil {
			log.Fatalf("failed to list tasks: %v", err)
		}

		for _, task := range tasks {
			fmt.Printf("Name: %s (Completed: %v)\n", task.Name, task.Completed)
		}
	},
}