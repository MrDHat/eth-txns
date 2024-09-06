package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockJSONRPCClient is a mock of the Client interface
type MockJSONRPCClient struct {
	mock.Mock
}

// MakeRequest mocks the MakeRequest method
func (m *MockJSONRPCClient) MakeRequest(method string, params []interface{}) (interface{}, error) {
	args := m.Called(method, params)
	return args.Get(0), args.Error(1)
}

// NewMockJSONRPCClient creates a new instance of MockJSONRPCClient
func NewMockJSONRPCClient() *MockJSONRPCClient {
	return &MockJSONRPCClient{}
}
