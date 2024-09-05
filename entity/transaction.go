package entity

type TransactionEntity struct {
	Hash        string
	From        string
	To          string
	Value       string
	BlockNumber int
}

type AddressTransactions map[string][]TransactionEntity // Takes too much memory but ok for now
