package servicenow

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("CreateTableRecord", func() {
	var (
		serviceNowAPIMock *APIRequestMock
		servicenow        ServiceNow
		config            Config
	)

	BeforeEach(func() {
		config = Config{
			User:     "user",
			Password: "password",
			BaseURL:  "https://example.service-now.com",
		}
		serviceNowAPIMock = new(APIRequestMock)
		servicenow = &serviceNow{
			config:     config,
			apiRequest: serviceNowAPIMock,
		}
	})

	It("should create a table record successfully", func() {
		tableName := "any_table_name"

		data := fakeMessageMock{
			Foo:           "foo",
			Bar:           "bar",
			CorrelationID: "any_correlation_id",
		}

		expectedChange := []byte(fmt.Sprintf(`
            {
                "sys_id": "abc123",
                "number": "CHG123",
                "foo": "%s",
                "bar": "%s",
                "correlation_id": "%s"
            }
        `, data.Foo, data.Bar, data.CorrelationID))

		serviceNowAPIMock.On("Post", "/api/now/table/any_table_name", mock.Anything).Return(201, expectedChange, nil)

		result, err := servicenow.CreateTableRecord(tableName, data)

		Expect(err).To(BeNil())
		Expect(result).To(MatchJSON(expectedChange))

		serviceNowAPIMock.AssertCalled(GinkgoT(), "Post", "/api/now/table/any_table_name", mock.Anything)
	})

	It("should return an error when failed to create a table record", func() {
		tableName := "any_table_name"

		data := interface{}(nil)

		mockError := fmt.Errorf("ServiceNow API error. Message: 'Invalid table name'. Detail: 'Table 'any_table_name' does not exist.'")

		serviceNowAPIMock.On("Post", "/api/now/table/any_table_name", mock.Anything).Return(401, nil, mockError)

		result, err := servicenow.CreateTableRecord(tableName, data)

		Expect(errors.Is(err, failedToCreateTableRecordError)).To(BeTrue(), "Expected failedToCreateTableRecordError")
		Expect(result).To(BeNil())

		serviceNowAPIMock.AssertCalled(GinkgoT(), "Post", "/api/now/table/any_table_name", mock.Anything)
	})
})
