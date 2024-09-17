package package_test

import (
	"testing"

	"github.com/mauricioromagnollo/servicenow-go/servicenow"
)

func TestImport(t *testing.T) {
	s := servicenow.NewServiceNow(servicenow.Config{
		User:     "admin",
		Password: "admin",
		BaseURL:  "https://dev00000.service-now.com",
	})

	if s == nil {
		t.Error("ServiceNow import failed")
	}
}
