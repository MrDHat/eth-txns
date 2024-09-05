package store

import (
	"github.com/mrdhat/eth-txns/entity"
	"github.com/mrdhat/eth-txns/errors"
)

type BlockStore interface {
	Save(block entity.BlockEntity) error
	GetLatest() (*entity.BlockEntity, error)
}

type blockStore struct {
	baseStore
	blocks []entity.BlockEntity
}

func (b *blockStore) Save(block entity.BlockEntity) error {
	switch b.storeType {
	case StoreTypeMemory:
		b.blocks = append(b.blocks, block)
		return nil
	default:
		return errors.ErrStoreTypeNotSupported
	}
}

func (b *blockStore) GetLatest() (*entity.BlockEntity, error) {
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
