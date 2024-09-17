package servicenow

import (
	"encoding/json"
	"errors"
	"fmt"
)

// CreateTableRecord creates a new record in the specified ServiceNow table.
// It takes the table name and data as input parameters and returns the response body as a byte array and an error, if any.
// The data parameter should be a struct or map that can be marshaled into JSON format.
// If successful, the response body will contain the result of the create operation.
// If an error occurs during the create operation, the function will return an error.
func (s *serviceNow) CreateTableRecord(tableName string, data interface{}) ([]byte, error) {
	urlPath := fmt.Sprintf("/api/now/table/%s", tableName)

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	_, r, err := s.apiRequest.Post(urlPath, body)
	if err != nil {
		return nil, errors.Join(failedToCreateTableRecordError, err)
	}

	return r, nil
}
