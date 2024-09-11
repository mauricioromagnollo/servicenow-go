package api

import (
	"net/http"

	"github.com/mauricioromagnollo/servicenow-go/servicenow"
)

type api struct {
	config     servicenow.Config
	httpClient *http.Client
}

// NewServiceNowApi creates a new instance of the ServiceNow API.
func NewAPI(config servicenow.Config) API {
	return &api{
		config:     config,
		httpClient: http.DefaultClient,
	}
}
