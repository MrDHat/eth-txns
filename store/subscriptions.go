package store

import "github.com/mrdhat/eth-txns/errors"

type AddressSubscription struct {
	Address  string
	IsActive bool
}

type AddressSubscriptionStore interface {
	Save(subscription AddressSubscription) error
	GetStatus(address string) (bool, error)
}

type addressSubscriptionStore struct {
	baseStore
	subscriptions map[string]AddressSubscription
}

func (s *addressSubscriptionStore) Save(subscription AddressSubscription) error {
	switch s.storeType {
	case StoreTypeMemory:
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
		subscriptions: make(map[string]AddressSubscription),
	}
}
