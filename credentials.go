package gocapi

import (
	"fmt"
	"os"
)

type Credentials struct {
	BaseUrl  string
	Username string
	Password string
	ClientId string
	Secret   string
}

func NewCredentials() Credentials {
	return Credentials{
		BaseUrl:  getBaseUrl(),
		Username: getUsername(),
		Password: getPassword(),
		ClientId: getClientId(),
		Secret:   getSecret(),
	}
}

func getEnv(envName string) string {
	value := os.Getenv(envName)

	if len(value) == 0 {
		msg := fmt.Sprintf("Env variable %v not set", envName)
		panic(msg)
	}

	return value
}

func getBaseUrl() string {
	return getEnv("GOCAPI_URL")
}

func getUsername() string {
	return getEnv("GOCAPI_USERNAME")
}

func getPassword() string {
	return getEnv("GOCAPI_PASSWORD")
}

func getClientId() string {
	return getEnv("GOCAPI_CLIENTID")
}

func getSecret() string {
	return getEnv("GOCAPI_SECRET")
}
