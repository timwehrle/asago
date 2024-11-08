package asago

import (
	"encoding/json"
	"os"
)

type Config struct {
	DefaultWorkspace string `json:"default_workspace"`
	DeleteToken      bool
}

var AppConfig Config

func LoadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&AppConfig)
}

func SaveConfig() error {
	data, err := json.MarshalIndent(AppConfig, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("config.json", data, 0644)
}