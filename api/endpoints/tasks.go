package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/timwehrle/asago/api"
	"github.com/timwehrle/asago/pkg/asago"
)

type TasksEndpoint struct {
	client *api.Client
}

func Tasks(client *api.Client) *TasksEndpoint {
	return &TasksEndpoint{client: client}
}

func (e *TasksEndpoint) ListTasks() ([]asago.Task, error) {
		err := asago.LoadConfig()
		if err != nil {
			return nil, err
		}
		
		endpoint := fmt.Sprintf("/tasks?assignee=me&workspace=%s", asago.AppConfig.DefaultWorkspace)
		req, err := e.client.New("GET", endpoint)
		if err != nil {
			return nil, err
		}

		resp, err := e.client.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
        	return nil, fmt.Errorf("failed to retrieve tasks, status: %v", resp.Status)
		}

		var tasksResponse struct {
			Data []asago.Task `json:"data"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&tasksResponse); err != nil {
			return nil, err
		}

		return tasksResponse.Data, nil
	}