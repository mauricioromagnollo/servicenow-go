package servicenow

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ServiceNow Suite")
}

type fakeMessageMock struct {
	Foo           string `json:"foo"`
	Bar           string `json:"bar"`
	CorrelationID string `json:"correlation_id"`
}

type APIRequestMock struct {
	APIRequest
	mock.Mock
}

func (a *APIRequestMock) Post(urlPath string, data []byte) (int, []byte, error) {
	args := a.Called(urlPath, data)
	if args.Get(1) == nil {
		return args.Int(0), nil, args.Error(2)
	}

	return args.Int(0), args.Get(1).([]byte), args.Error(2)
}

func (a *APIRequestMock) Put(urlPath string, data []byte) (int, []byte, error) {
	args := a.Called(urlPath, data)
	if args.Get(1) == nil {
		return args.Int(0), nil, args.Error(2)
	}

	return args.Int(0), args.Get(1).([]byte), args.Error(2)
}

func (a *APIRequestMock) Patch(urlPath string, data []byte) (int, []byte, error) {
	args := a.Called(urlPath, data)
	if args.Get(1) == nil {
		return args.Int(0), nil, args.Error(2)
	}

	return args.Int(0), args.Get(1).([]byte), args.Error(2)
}

func (a *APIRequestMock) Delete(urlPath string) (int, []byte, error) {
	args := a.Called(urlPath)
	if args.Get(1) == nil {
		return args.Int(0), nil, args.Error(2)
	}

	return args.Int(0), args.Get(1).([]byte), args.Error(2)
}

func (a *APIRequestMock) Get(urlPath string) (int, []byte, error) {
	args := a.Called(urlPath)
	if args.Get(1) == nil {
		return args.Int(0), nil, args.Error(2)
	}

	return args.Int(0), args.Get(1).([]byte), args.Error(2)
}
