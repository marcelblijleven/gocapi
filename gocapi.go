package gocapi

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

// Client is used to interact with the ocapi endpoints
type Client struct {
	BaseURL    *url.URL
	Username   string
	Password   string
	ClientID   string
	Secret     string
	httpClient http.Client

	Authentication *Authentication
}

// NewClient creates a new Client with credentials stored in env variables
func NewClient(httpClient http.Client) (*Client, error) {
	creds := NewCredentials()

	if len(creds.BaseUrl) == 0 {
		return nil, errors.New("Base url is required")
	}

	u, err := url.Parse(creds.BaseUrl)

	if err != nil {
		return nil, err
	}

	c := &Client{
		BaseURL:    u,
		Username:   creds.Username,
		Password:   creds.Password,
		ClientID:   creds.ClientId,
		Secret:     creds.Secret,
		httpClient: httpClient,
	}

	c.Authentication = &Authentication{Client: c}
	return c, nil
}

// CreateRequest uses the httpClient in Client to create a new request
func (c Client) CreateRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	relPath, err := url.Parse(endpoint)

	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(relPath)

	req, err := http.NewRequest(method, u.String(), body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.Authentication.Token)
	req.Header.Add("Content-Type", "application/json")

	return req, err
}
