package servicenow

// Config represents the configuration for the ServiceNow client.
type Config struct {
	User     string // User is the username for authentication.
	Password string // Password is the password for authentication.
	BaseURL  string // BaseURL is the base URL of the ServiceNow instance.
}
