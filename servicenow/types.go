package servicenow

import "encoding/json"

// Config represents the configuration for the ServiceNow client.
type Config struct {
	User     string // User is the username for authentication.
	Password string // Password is the password for authentication.
	BaseURL  string // BaseURL is the base URL of the ServiceNow instance.
}

type servicenowApiResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *apiError       `json:"error"`
	Status *string         `json:"status"`
}

type apiError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
