package asago

type User struct {
	Gid        string      `json:"gid"`
	Email      string      `json:"email"`
	Name       string      `json:"name"`
	Workspaces []Workspace `json:"workspaces"`
}

type Workspace struct {
	Gid          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type UserResponse struct {
	Data User `json:"data"`
}