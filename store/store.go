package store

type StoreType string

const (
	StoreTypeMemory StoreType = "memory"
)

type baseStore struct {
	storeType StoreType
}
