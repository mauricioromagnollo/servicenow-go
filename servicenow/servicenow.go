package servicenow

// ServiceNow is an interface that defines the methods that are required to interact with ServiceNow.
type ServiceNow interface {
	// CreateTableRecord creates a change in ServiceNow.
	CreateTableRecord(tableName string, data interface{}) ([]byte, error)
}
