package node

import (
	"time"

	"github.com/mrdhat/eth-txns/helpers"
	"github.com/mrdhat/eth-txns/jsonrpc"
	"github.com/mrdhat/eth-txns/store"
)

type Listener interface {
	Start(stop <-chan bool) error
}

type listener struct {
	currentBlockDuration int

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

	for {
		select {
		case <-stop:
			return nil
		default:
			blockNumberToRead := latestBlock.Number
			// read the latest block
			block, err := l.readBlock(blockNumberToRead)
			if err != nil {
				return err
			}

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

			// update the latest block
			currBlockNumber, err := helpers.ConvertHexToDecimal(block.Number)
			if err != nil {
				return err
			}
			blockNumberToRead = helpers.ConvertPositiveDecimalToHex(currBlockNumber + 1)

			// sleep for the duration of the block
			time.Sleep(time.Duration(l.currentBlockDuration) * time.Second)
		}
	}

}

func (l *listener) getLatestBlock() (*blockchainBlock, error) {
	block, err := l.rpcClient.MakeRequest("eth_blockNumber", []interface{}{})
	if err != nil {
		return nil, err
	}

	blockData := block.(blockchainBlock)

	return &blockData, nil
}

func (l *listener) readBlock(number string) (*blockchainBlock, error) {
	block, err := l.rpcClient.MakeRequest("eth_getBlockByNumber", []interface{}{number, true})
	if err != nil {
		return nil, nil
	}

	blockData := block.(blockchainBlock)

	return &blockData, nil
}

func NewListener(currentBlockDuration int, jsonRPCClient jsonrpc.Client, blockStore store.BlockStore, transactionStore store.TransactionStore) Listener {
	return &listener{
		rpcClient:        jsonRPCClient,
		blockStore:       blockStore,
		transactionStore: transactionStore,
	}
}
