package api

import (
	"github.com/mrdhat/eth-txns/entity"
	"github.com/mrdhat/eth-txns/store"
)

type Parser interface {
	// last parsed block
	GetCurrentBlock() int
	// add address to observer
	Subscribe(address string) error
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []entity.TransactionEntity
}

type parser struct {
	blockStore       store.BlockStore
	addressStore     store.AddressSubscriptionStore
	transactionStore store.TransactionStore
}

func (p *parser) GetCurrentBlock() int {
	block, err := p.blockStore.GetLatest()
	if err != nil {
		return 0
	}
	return block.Number
}

func (p *parser) Subscribe(address string) error {
	subscription := entity.AddressSubscription{
		Address:  address,
		IsActive: true,
	}
	return p.addressStore.Save(subscription)
}

func (p *parser) GetTransactions(address string) []entity.TransactionEntity {
	transactions, err := p.transactionStore.GetAllByAddress(address)
	if err != nil {
		return nil
	}
	return transactions
}

func NewParser(blockStore store.BlockStore, addressStore store.AddressSubscriptionStore, transactionStore store.TransactionStore) Parser {
	return &parser{
		blockStore:       blockStore,
		addressStore:     addressStore,
		transactionStore: transactionStore,
	}
}
