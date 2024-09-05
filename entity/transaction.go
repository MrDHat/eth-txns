package entity

type TransactionEntity struct {
	Hash        string
	From        string
	To          string
	Value       string
	BlockNumber int
}
