package store

import (
	"testing"

	"github.com/mrdhat/eth-txns/entity"
	"github.com/mrdhat/eth-txns/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewTransactionStore(t *testing.T) {
	addressSubscriptionStore := NewAddressSubscriptionStore(StoreTypeMemory)
	store := NewTransactionStore(StoreTypeMemory, addressSubscriptionStore)
	assert.NotNil(t, store)
	assert.IsType(t, &transactionStore{}, store)
}

func TestTransactionStore_Save(t *testing.T) {
	addressSubscriptionStore := NewAddressSubscriptionStore(StoreTypeMemory)
	store := NewTransactionStore(StoreTypeMemory, addressSubscriptionStore).(*transactionStore)

	transactions := []entity.TransactionEntity{
		{From: "0x1", To: "0x2", Value: "100"},
		{From: "0x2", To: "0x3", Value: "200"},
	}

	// subscribe the addresses
	addressSubscriptionStore.Save(entity.AddressSubscription{Address: "0x1", IsActive: true})
	addressSubscriptionStore.Save(entity.AddressSubscription{Address: "0x2", IsActive: true})

	err := store.Save(transactions)
	assert.NoError(t, err)

	assert.Len(t, store.transactions["0x1"], 1)
	assert.Len(t, store.transactions["0x2"], 2)
}

func TestTransactionStore_GetAllByAddress(t *testing.T) {
	addressSubscriptionStore := NewAddressSubscriptionStore(StoreTypeMemory)
	store := NewTransactionStore(StoreTypeMemory, addressSubscriptionStore).(*transactionStore)

	transactions := []entity.TransactionEntity{
		{From: "0x1", To: "0x2", Value: "100"},
		{From: "0x4", To: "0x3", Value: "200"},
		{From: "0x1", To: "0x3", Value: "300"},
	}

	// subscribe the addresses
	addressSubscriptionStore.Save(entity.AddressSubscription{Address: "0x1", IsActive: true})
	addressSubscriptionStore.Save(entity.AddressSubscription{Address: "0x2", IsActive: true})
	addressSubscriptionStore.Save(entity.AddressSubscription{Address: "0x3", IsActive: true})

	_ = store.Save(transactions)

	tests := []struct {
		name    string
		address string
		want    int
	}{
		{"Address with 2 transactions", "0x1", 2},
		{"Address with 1 transaction", "0x2", 1},
		{"Address with no transactions", "0x5", 0},
		{"Address with no subscription", "0x4", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := store.GetAllByAddress(tt.address)
			assert.Len(t, result, tt.want)
		})
	}
}

func TestTransactionStore_SaveUnsupportedStore(t *testing.T) {
	addressSubscriptionStore := NewAddressSubscriptionStore(StoreTypeMemory)
	store := NewTransactionStore(StoreType("unsupported"), addressSubscriptionStore)

	transactions := []entity.TransactionEntity{
		{From: "0x1", To: "0x2", Value: "100"},
	}

	err := store.Save(transactions)
	assert.Error(t, err)
	assert.Equal(t, errors.ErrStoreTypeNotSupported, err)
}

func TestTransactionStore_GetAllByAddressUnsupportedStore(t *testing.T) {
	addressSubscriptionStore := NewAddressSubscriptionStore(StoreTypeMemory)
	store := NewTransactionStore(StoreType("unsupported"), addressSubscriptionStore)

	result := store.GetAllByAddress("0x1")
	assert.Nil(t, result)
}
