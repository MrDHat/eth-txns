package mocks

import (
	"github.com/mrdhat/eth-txns/entity"
	"github.com/stretchr/testify/mock"
)

type MockBlockStore struct {
	mock.Mock
}

func (m *MockBlockStore) Save(block entity.BlockEntity) error {
	args := m.Called(block)
	return args.Error(0)
}

func (m *MockBlockStore) GetLatest() (*entity.BlockEntity, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.BlockEntity), args.Error(1)
}

type MockAddressSubscriptionStore struct {
	mock.Mock
}

func (m *MockAddressSubscriptionStore) Save(subscription entity.AddressSubscription) error {
	args := m.Called(subscription)
	return args.Error(0)
}

func (m *MockAddressSubscriptionStore) GetStatus(address string) (bool, error) {
	args := m.Called(address)
	return args.Bool(0), args.Error(1)
}

type MockTransactionStore struct {
	mock.Mock
}

func (m *MockTransactionStore) Save(transactions []entity.TransactionEntity) error {
	args := m.Called(transactions)
	return args.Error(0)
}

func (m *MockTransactionStore) GetAllByAddress(address string) []entity.TransactionEntity {
	args := m.Called(address)
	return args.Get(0).([]entity.TransactionEntity)
}
