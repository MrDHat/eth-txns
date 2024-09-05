package store

import (
	"github.com/mrdhat/eth-txns/entity"
	"github.com/mrdhat/eth-txns/errors"
	"github.com/mrdhat/eth-txns/logger"
)

type TransactionStore interface {
	Save(transactions []entity.TransactionEntity) error
	GetAllByAddress(address string) []entity.TransactionEntity
	GetAll() entity.AddressTransactions
}

type transactionStore struct {
	baseStore
	addressSubscriptionStore AddressSubscriptionStore
	transactions             entity.AddressTransactions
}

func (t *transactionStore) Save(transactions []entity.TransactionEntity) error {
	switch t.storeType {
	case StoreTypeMemory:
		for _, transaction := range transactions {
			isFromSubscribed, err := t.addressSubscriptionStore.GetStatus(transaction.From)
			if err != nil {
				return err
			}
			isToSubscribed, err := t.addressSubscriptionStore.GetStatus(transaction.To)
			if err != nil {
				return err
			}

			if isFromSubscribed {
				logger.Log("Saving a transaction from: ", transaction.From)
				t.transactions[transaction.From] = append(t.transactions[transaction.From], transaction)
			}
			if isToSubscribed {
				logger.Log("Saving a transaction to: ", transaction.To)
				t.transactions[transaction.To] = append(t.transactions[transaction.To], transaction)
			}
		}
		return nil
	default:
		return errors.ErrStoreTypeNotSupported
	}
}

func (t *transactionStore) GetAllByAddress(address string) []entity.TransactionEntity {
	switch t.storeType {
	case StoreTypeMemory:
		txns, ok := t.transactions[address]
		if !ok {
			return nil
		}
		return txns
	default:
		return nil
	}
}

func (t *transactionStore) GetAll() entity.AddressTransactions {
	switch t.storeType {
	case StoreTypeMemory:
		return t.transactions
	default:
		return nil
	}
}

func NewTransactionStore(storeType StoreType, addressSubscriptionStore AddressSubscriptionStore) TransactionStore {
	return &transactionStore{
		baseStore:                baseStore{storeType: storeType},
		addressSubscriptionStore: addressSubscriptionStore,
		transactions:             make(entity.AddressTransactions),
	}
}
