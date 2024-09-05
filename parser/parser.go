package parser

import (
	"github.com/mrdhat/eth-txns/transaction"
)

type Parser interface {
	// last parsed block
	GetCurrentBlock() int
	// add address to observer
	Subscribe(address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []transaction.Transaction
}

type parser struct {
}

func (p *parser) GetCurrentBlock() int {
	return 0
}

func (p *parser) Subscribe(address string) bool {
	return false
}

func (p *parser) GetTransactions(address string) []transaction.Transaction {
	return nil
}

func NewParser() Parser {
	return &parser{}
}
