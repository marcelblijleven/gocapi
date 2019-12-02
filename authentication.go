package gocapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// Authentication service
type Authentication struct {
	Client *Client
	Token  string
}

// AuthenticationResponse from ocapi oauth request
type AuthenticationResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// Authenticate the gocapi client
func (a Authentication) Authenticate() error {
	endpoint := fmt.Sprintf("/dw/oauth2/access_token?client_id=%v", a.Client.ClientID)
	data := url.Values{}
	data.Set("grant_type", "urn:demandware:params:oauth:grant-type:client-id:dwsid:dwsecuretoken")

	req, err := a.Client.CreateRequest("POST", endpoint, strings.NewReader(data.Encode()))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(a.Client.Username, fmt.Sprintf("%v:%v", a.Client.Password, a.Client.Secret))

	res, err := a.Client.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	authResponse := &AuthenticationResponse{}
	json.NewDecoder(res.Body).Decode(authResponse)

	a.Client.Authentication.Token = fmt.Sprintf("Bearer %v", authResponse.AccessToken)

	return nil
}
