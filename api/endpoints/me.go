package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/timwehrle/asago/api"
	"github.com/timwehrle/asago/pkg/asago"
)

type MeEndpoint struct {
	client *api.Client
}

func Me(client *api.Client) *MeEndpoint {
	return &MeEndpoint{client: client}
}

func (e *MeEndpoint) Get() (*asago.User, error) {
	req, err := e.client.New("GET", "/users/me")
	if err != nil {
		return nil, err
	}

	resp, err := e.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve user info, status: %v", resp.Status)
	}

	var userResponse struct {
		Data asago.User `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return nil, err
	}

	return &userResponse.Data, nil
}