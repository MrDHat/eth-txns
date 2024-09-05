package store

import "github.com/mrdhat/eth-txns/entity"

type TransactionStore interface {
	SaveTransactions(transactions []entity.TransactionEntity) error
	GetAllTransactionsByAddress(address string) []entity.TransactionEntity
}
