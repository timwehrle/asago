package asago

import (
	"github.com/zalando/go-keyring"
)

const (
	service  = "asago"
	username = "user"
)

func GetToken() (string, error) {
	return keyring.Get(service, username)
}

func SetToken(token string) error {
	return keyring.Set(service, username, token)
}

func DeleteToken() error {
	return keyring.Delete(service, username)
}