package gocapi

import (
	"errors"
	"gocapi/util"
)

type Client struct {
	BaseUrl  string
	Username string
	Password string
	ClientId string
	Secret   string
}

func NewClient() (*Client, error) {
	c := util.NewCredentials()

	if len(c.BaseUrl) == 0 {
		return nil, errors.New("Base url is required")
	}

	return &Client{
		BaseUrl:  c.BaseUrl,
		Username: c.Username,
		Password: c.Password,
		ClientId: c.ClientId,
		Secret:   c.Secret,
	}, nil
}
