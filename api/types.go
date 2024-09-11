package api

import "encoding/json"

type servicenowApiResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *apiError       `json:"error"`
	Status *string         `json:"status"`
}

type apiError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
