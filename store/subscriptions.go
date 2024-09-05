package store

import (
	"github.com/mrdhat/eth-txns/entity"
	"github.com/mrdhat/eth-txns/errors"
	"github.com/mrdhat/eth-txns/logger"
)

type AddressSubscriptionStore interface {
	Save(subscription entity.AddressSubscription) error
	GetStatus(address string) (bool, error)
}

type addressSubscriptionStore struct {
	baseStore
	subscriptions entity.AddressSubscriptionEntity
}

func (s *addressSubscriptionStore) Save(subscription entity.AddressSubscription) error {
	switch s.storeType {
	case StoreTypeMemory:
		logger.Log("Saving subscription: ", subscription)
		s.subscriptions[subscription.Address] = subscription
		return nil
	default:
		return errors.ErrStoreTypeNotSupported
	}
}

func (s *addressSubscriptionStore) GetStatus(address string) (bool, error) {
	switch s.storeType {
	case StoreTypeMemory:
		subscription, ok := s.subscriptions[address]
		if !ok {
			return false, nil
		}
		return subscription.IsActive, nil
	default:
		return false, errors.ErrStoreTypeNotSupported
	}
}

func NewAddressSubscriptionStore(storeType StoreType) AddressSubscriptionStore {
	return &addressSubscriptionStore{
		baseStore: baseStore{
			storeType: storeType,
		},
		subscriptions: make(entity.AddressSubscriptionEntity),
	}
}
