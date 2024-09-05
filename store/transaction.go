package store

import "github.com/mrdhat/eth-txns/entity"

type TransactionStore interface {
	Save(transactions []entity.TransactionEntity) error
	GetAllByAddress(address string) []entity.TransactionEntity
}
