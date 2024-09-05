package node

import (
	"time"

	"github.com/mrdhat/eth-txns/entity"
	"github.com/mrdhat/eth-txns/helpers"
)

type blockchainTransaction struct {
	Hash      string `json:"hash"`
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

type blockchainBlock struct {
	Number       string                  `json:"number"`
	Hash         string                  `json:"hash"`
	Timestamp    string                  `json:"timestamp"`
	Transactions []blockchainTransaction `json:"transactions"`
}

func (b *blockchainBlock) ToStoreEntities() (entity.BlockEntity, []entity.TransactionEntity, error) {
	blockNumber, err := helpers.ConvertHexToDecimal(b.Number)
	if err != nil {
		return entity.BlockEntity{}, nil, err
	}

	blockTimestamp, err := helpers.ConvertHexToDecimal(b.Timestamp)
	if err != nil {
		return entity.BlockEntity{}, nil, err
	}

	blockEntity := entity.BlockEntity{
		Number:    blockNumber,
		Hash:      b.Hash,
		Timestamp: time.Unix(int64(blockTimestamp), 0),
	}

	transactionEntities := make([]entity.TransactionEntity, len(b.Transactions))
	for i, transaction := range b.Transactions {
		transactionEntities[i] = entity.TransactionEntity{
			Hash:        transaction.Hash,
			From:        transaction.From,
			To:          transaction.To,
			Value:       transaction.Value,
			BlockNumber: blockNumber,
		}
	}

	return blockEntity, transactionEntities, nil
}
