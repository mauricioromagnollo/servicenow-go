package servicenow

import (
	"fmt"
)

var (
	failedToCreateTableRecordError error = fmt.Errorf("Failed to create table record")
)

func genericApiError(err *apiError) error {
	return fmt.Errorf("ServiceNow API error. Message: '%s'. Detail: '%s'", err.Message, err.Detail)
}
