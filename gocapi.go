package gocapi

import (
	"errors"
)

// Client is used to interact with the ocapi endpoints
type Client struct {
	BaseURL  string
	Username string
	Password string
	ClientID string
	Secret   string
}

// NewClient creates a new Client with credentials stored in env variables
func NewClient() (*Client, error) {
	c := NewCredentials()

	if len(c.BaseUrl) == 0 {
		return nil, errors.New("Base url is required")
	}

	return &Client{
		BaseURL:  c.BaseUrl,
		Username: c.Username,
		Password: c.Password,
		ClientID: c.ClientId,
		Secret:   c.Secret,
	}, nil
}
