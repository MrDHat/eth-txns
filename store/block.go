package store

import (
	"time"

	"github.com/mrdhat/eth-txns/errors"
)

type BlockStore interface {
	Save(block Block) error
	GetLatest() (*Block, error)
}

type Block struct {
	Number    int
	Hash      string
	Timestamp time.Time
}

type blockStore struct {
	baseStore
	blocks []Block
}

func (b *blockStore) Save(block Block) error {
	switch b.storeType {
	case StoreTypeMemory:
		b.blocks = append(b.blocks, block)
		return nil
	default:
		return errors.ErrStoreTypeNotSupported
	}
}

func (b *blockStore) GetLatest() (*Block, error) {
	switch b.storeType {
	case StoreTypeMemory:
		if len(b.blocks) == 0 {
			return nil, nil
		}
		return &b.blocks[len(b.blocks)-1], nil
	default:
		return nil, errors.ErrStoreTypeNotSupported
	}
}

func NewBlockStore(storeType StoreType) BlockStore {
	return &blockStore{
		baseStore: baseStore{storeType: storeType},
	}
}
