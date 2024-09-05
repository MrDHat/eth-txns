package mocks

import "github.com/stretchr/testify/mock"

type MockRPCClient struct {
	mock.Mock
}

func (m *MockRPCClient) MakeRequest(method string, params []interface{}) (interface{}, error) {
	args := m.Called(method, params)
	return args.Get(0), args.Error(1)
}
