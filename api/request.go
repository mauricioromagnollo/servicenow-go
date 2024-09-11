package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

	basicAuth, err := getBasicWithEncodedCredentials(a.config.User, a.config.Password)
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
