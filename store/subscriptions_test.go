package store_test

import (
	"testing"

	"github.com/mrdhat/eth-txns/entity"
	"github.com/mrdhat/eth-txns/errors"
	"github.com/mrdhat/eth-txns/store"
	"github.com/stretchr/testify/assert"
)

func TestNewAddressSubscriptionStore(t *testing.T) {
	s := store.NewAddressSubscriptionStore(store.StoreTypeMemory)
	assert.NotNil(t, s)
}

func TestAddressSubscriptionStore_Save(t *testing.T) {
	s := store.NewAddressSubscriptionStore(store.StoreTypeMemory)

	tests := []struct {
		name         string
		subscription entity.AddressSubscription
		wantErr      error
	}{
		{
			name: "Save valid subscription",
			subscription: entity.AddressSubscription{
				Address:  "0x1234567890123456789012345678901234567890",
				IsActive: true,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Save(tt.subscription)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestAddressSubscriptionStore_GetStatus(t *testing.T) {
	s := store.NewAddressSubscriptionStore(store.StoreTypeMemory)

	// Save a test subscription
	testSubscription := entity.AddressSubscription{
		Address:  "0x1234567890123456789012345678901234567890",
		IsActive: true,
	}
	err := s.Save(testSubscription)
	assert.NoError(t, err)

	tests := []struct {
		name       string
		address    string
		wantActive bool
		wantErr    error
	}{
		{
			name:       "Get status of existing subscription",
			address:    "0x1234567890123456789012345678901234567890",
			wantActive: true,
			wantErr:    nil,
		},
		{
			name:       "Get status of non-existing subscription",
			address:    "0x0987654321098765432109876543210987654321",
			wantActive: false,
			wantErr:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			active, err := s.GetStatus(tt.address)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantActive, active)
		})
	}
}

func TestAddressSubscriptionStore_UnsupportedStoreType(t *testing.T) {
	s := store.NewAddressSubscriptionStore(store.StoreType("unsupported"))

	t.Run("Save with unsupported store type", func(t *testing.T) {
		err := s.Save(entity.AddressSubscription{})
		assert.Equal(t, errors.ErrStoreTypeNotSupported, err)
	})

	t.Run("GetStatus with unsupported store type", func(t *testing.T) {
		_, err := s.GetStatus("0x1234567890123456789012345678901234567890")
		assert.Equal(t, errors.ErrStoreTypeNotSupported, err)
	})
}
