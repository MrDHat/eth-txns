package node

import (
	"fmt"
	"time"

	"github.com/mrdhat/eth-txns/errors"
	"github.com/mrdhat/eth-txns/helpers"
	"github.com/mrdhat/eth-txns/jsonrpc"
	"github.com/mrdhat/eth-txns/logger"
	"github.com/mrdhat/eth-txns/store"
)

type Listener interface {
	Start(stop <-chan bool) error
}

type listener struct {
	currentBlockDurationInSeconds int

	rpcClient jsonrpc.Client

	blockStore       store.BlockStore
	transactionStore store.TransactionStore
}

func (l *listener) Start(stop <-chan bool) error {
	// read the latest block
	latestBlock, err := l.getLatestBlock()
	if err != nil {
		return err
	}
	if latestBlock == nil {
		return errors.ErrFailedToGetLatestBlock
	}

	blockNumberToRead := *latestBlock
	for {
		select {
		case <-stop:
			return nil
		default:
			// read the latest block
			block, err := l.readBlock(blockNumberToRead)
			if err != nil {
				return err
			}

			if block == nil {
				logger.Log("empty block")
			} else {
				// Save the entities to the database
				blockEntity, transactionEntities, err := block.ToStoreEntities()
				if err != nil {
					return err
				}

				err = l.blockStore.Save(blockEntity)
				if err != nil {
					return err
				}

				err = l.transactionStore.Save(transactionEntities)
				if err != nil {
					return err
				}

				currBlockNumber, err := helpers.ConvertHexToDecimal(blockNumberToRead)
				if err != nil {
					return err
				}

				// update the latest block
				blockNumberToRead = helpers.ConvertPositiveDecimalToHex(currBlockNumber + 1)

			}

			// sleep for the duration of the block
			time.Sleep(time.Duration(l.currentBlockDurationInSeconds) * time.Second)
		}
	}

}

func (l *listener) getLatestBlock() (*string, error) {
	block, err := l.rpcClient.MakeRequest("eth_blockNumber", []interface{}{})
	if err != nil {
		return nil, err
	}

	blockData := block.(string)

	return &blockData, nil
}

func (l *listener) readBlock(number string) (*blockchainBlock, error) {
	block, err := l.rpcClient.MakeRequest("eth_getBlockByNumber", []interface{}{number, true})
	if err != nil {
		return nil, err
	}

	if block == nil {
		// if somehow the block is empty
		return nil, nil
	}

	blockDataRes := block.(map[string]interface{})

	transactions := blockDataRes["transactions"].([]interface{})
	blockchainTransactions := make([]blockchainTransaction, len(transactions))

	logger.Log(fmt.Sprintf("Found %v transactions for block %v: ", len(transactions), blockDataRes["number"]))

	for i, transaction := range transactions {
		transactionMap := transaction.(map[string]interface{})
		blockchainTransactions[i] = blockchainTransaction{
			Hash:      transactionMap["hash"].(string),
			From:      transactionMap["from"].(string),
			To:        transactionMap["to"].(string),
			Value:     transactionMap["value"].(string),
			Timestamp: blockDataRes["timestamp"].(string),
		}
	}

	return &blockchainBlock{
		Number:       blockDataRes["number"].(string),
		Hash:         blockDataRes["hash"].(string),
		Timestamp:    blockDataRes["timestamp"].(string),
		Transactions: blockchainTransactions,
	}, nil
}

func NewListener(currentBlockDurationInSeconds int, jsonRPCClient jsonrpc.Client, blockStore store.BlockStore, transactionStore store.TransactionStore) Listener {
	return &listener{
		currentBlockDurationInSeconds: currentBlockDurationInSeconds,
		rpcClient:                     jsonRPCClient,
		blockStore:                    blockStore,
		transactionStore:              transactionStore,
	}
}
