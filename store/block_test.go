package store_test

import (
	"testing"
	"time"

	"github.com/mrdhat/eth-txns/entity"
	"github.com/mrdhat/eth-txns/errors"
	"github.com/mrdhat/eth-txns/store"
	"github.com/stretchr/testify/assert"
)

func TestNewBlockStore(t *testing.T) {
	bs := store.NewBlockStore(store.StoreTypeMemory)
	assert.NotNil(t, bs, "NewBlockStore should return a non-nil BlockStore")
}

func TestSaveBlock(t *testing.T) {
	bs := store.NewBlockStore(store.StoreTypeMemory)

	block := entity.BlockEntity{
		Number:    1,
		Hash:      "0x123",
		Timestamp: time.Now(),
	}

	err := bs.Save(block)
	assert.NoError(t, err, "SaveBlock should not return an error for StoreTypeMemory")

	// Test unsupported store type
	unsupportedBS := store.NewBlockStore(store.StoreType("unsupported"))
	err = unsupportedBS.Save(block)
	assert.Equal(t, errors.ErrStoreTypeNotSupported, err, "SaveBlock should return ErrStoreTypeNotSupported for unsupported store type")
}

func TestGetLatestBlock(t *testing.T) {
	bs := store.NewBlockStore(store.StoreTypeMemory)

	// Test empty store
	latestBlock, err := bs.GetLatest()
	assert.NoError(t, err, "GetLatestBlock should not return an error for empty store")
	assert.Nil(t, latestBlock, "GetLatestBlock should return nil for empty store")

	// Add blocks and test
	blocks := []entity.BlockEntity{
		{Number: 1, Hash: "0x123", Timestamp: time.Now()},
		{Number: 2, Hash: "0x456", Timestamp: time.Now()},
		{Number: 3, Hash: "0x789", Timestamp: time.Now()},
	}

	for _, block := range blocks {
		err := bs.Save(block)
		assert.NoError(t, err, "SaveBlock should not return an error")
	}

	latestBlock, err = bs.GetLatest()
	assert.NoError(t, err, "GetLatestBlock should not return an error")
	assert.Equal(t, 3, latestBlock.Number, "GetLatestBlock should return the number of the latest block")

	// Test unsupported store type
	unsupportedBS := store.NewBlockStore(store.StoreType("unsupported"))
	_, err = unsupportedBS.GetLatest()
	assert.Equal(t, errors.ErrStoreTypeNotSupported, err, "GetLatestBlock should return ErrStoreTypeNotSupported for unsupported store type")
}
