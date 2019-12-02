package gocapi

import (
	"encoding/json"
	"fmt"
	"time"
)

// CodeVersionGetResult ocapi get code versions result
type CodeVersionGetResult struct {
	OcapiVersion string        `json:"_v"`
	ResultType   string        `json:"_type"`
	Count        int           `json:"count"`
	Data         []CodeVersion `json:"data"`
	Total        int           `json:"total"`
}

// CodeVersion ocapi result
type CodeVersion struct {
	Type                 string    `json:"_type"`
	ResourceState        string    `json:"_resource_state"`
	ActivationTime       time.Time `json:"activation_time"`
	Active               bool      `json:"active"`
	CompatibilityMode    string    `json:"compatibility_mode"`
	ID                   string    `json:"id"`
	LastModificationTime time.Time `json:"last_modification_time"`
	Rollback             bool      `json:"rollback"`
	TotalSize            string    `json:"total_size"`
	WebDavURL            string    `json:"web_dav_url"`
}

// CodeVersionService code versions
type CodeVersionService struct {
	Client *Client
}

// Get code versions
func (c CodeVersionService) Get() (CodeVersionGetResult, error) {
	endpoint := "/s/-/dw/data/v19_10/code_versions"
	var r CodeVersionGetResult

	req, err := c.Client.CreateRequest("GET", endpoint, nil)

	if err != nil {
		return r, err
	}

	res, err := c.Client.httpClient.Do(req)

	if err != nil {
		return r, err
	}

	fmt.Println(res.StatusCode)

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&r)

	return r, nil
}
