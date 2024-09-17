package servicenow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mauricioromagnollo/servicenow-go/helper"
)

// Request is the interface that wraps the basic methods to interact with the ServiceNow API.
type APIRequest interface {
	Get(urlPath string) (int, []byte, error)
	Post(urlPath string, data []byte) (int, []byte, error)
	Put(urlPath string, data []byte) (int, []byte, error)
	Delete(urlPath string) (int, []byte, error)
	Patch(urlPath string, data []byte) (int, []byte, error)
}

type api struct {
	config     Config
	httpClient *http.Client
}

// APIRequest creates a new instance of the ServiceNow API.
func NewAPIRequest(config Config) APIRequest {
	return &api{
		config:     config,
		httpClient: http.DefaultClient,
	}
}

// Put sends a PUT request to the ServiceNow API.
func (a *api) Put(urlPath string, data []byte) (int, []byte, error) {
	return a.request("PUT", urlPath, data)
}

// Post sends a POST request to the ServiceNow API.
func (a *api) Post(urlPath string, data []byte) (int, []byte, error) {
	return a.request("POST", urlPath, data)
}

// Patch sends a PATCH request to the ServiceNow API.
func (a *api) Patch(urlPath string, data []byte) (int, []byte, error) {
	return a.request("PATCH", urlPath, data)
}

// Delete sends a DELETE request to the ServiceNow API.
func (a *api) Delete(urlPath string) (int, []byte, error) {
	return a.request("DELETE", urlPath, nil)
}

// Get sends a GET request to the ServiceNow API.
func (a *api) Get(urlPath string) (int, []byte, error) {
	return a.request("GET", urlPath, nil)
}

// request is an internal helper function for making HTTP requests to the API.
// This method returns the HTTP status code, the response body, and an error.
func (a *api) request(method, urlPath string, body []byte) (int, []byte, error) {
	if urlPath[0] != '/' {
		return 0, nil, fmt.Errorf("The param 'urlPath' must start with a slash")
	}

	url := fmt.Sprintf("%s%s", a.config.BaseURL, urlPath)

	var bodyReader io.Reader = nil
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return 0, nil, err
	}

	basicAuth, err := helper.GetBasicWithEncodedCredentials(a.config.User, a.config.Password)
	if err != nil {
		return 0, nil, err
	}

	req.Header.Add("Authorization", basicAuth)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := a.httpClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}

	response := servicenowApiResponse{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return res.StatusCode, b, err
	}

	if response.Error != nil {
		return res.StatusCode, b, fmt.Errorf("ServiceNow API error. Message: '%s'. Detail: '%s'", response.Error.Message, response.Error.Detail)
	}

	return res.StatusCode, response.Result, nil
}
