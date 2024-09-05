package parser

import "github.com/mrdhat/eth-txns/entity"

type Parser interface {
	// last parsed block
	GetCurrentBlock() int
	// add address to observer
	Subscribe(address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []entity.TransactionEntity
}

type parser struct {
}

func (p *parser) GetCurrentBlock() int {
	return 0
}

func (p *parser) Subscribe(address string) bool {
	return false
}

func (p *parser) GetTransactions(address string) []entity.TransactionEntity {
	return nil
}

func NewParser() Parser {
	return &parser{}
}
