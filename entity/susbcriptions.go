package entity

type AddressSubscription struct {
	Address  string
	IsActive bool
}

type AddressSubscriptionEntity map[string]AddressSubscription
