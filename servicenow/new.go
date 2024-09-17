package servicenow

type serviceNow struct {
	ServiceNow
	config     Config
	apiRequest apiRequest
}

// NewServiceNow creates a new ServiceNow client instance.
// It requires a servicenow.Config struct as input.
func NewServiceNow(config Config) ServiceNow {
	apiRequest := newAPIRequest(config)

	return &serviceNow{
		config:     config,
		apiRequest: apiRequest,
	}
}
