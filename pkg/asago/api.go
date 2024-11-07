package asago

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zalando/go-keyring"
)

const (
	service = "asago"
	username = "user"
)

func GetToken() (string, error) {
	token, err := keyring.Get(service, username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func FetchUserInfo() {
	token, err := GetToken()
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://app.asana.com/api/1.0/users/me", nil)
	req.Header.Add("Authorization", "Bearer " + token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to retrieve user info. Status: %v\n", resp.Status)
		return
	}

	var userResponse UserResponse
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	if err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}
	
	fmt.Printf("User Info:\nGID: %s\nName: %s\nEmail: %s\n", userResponse.Data.Gid, userResponse.Data.Name, userResponse.Data.Email)
	for _, workspace := range userResponse.Data.Workspaces {
		fmt.Printf("Workspace: %s (ID: %s)\n", workspace.Name, workspace.Gid)
	}
}