package db

import "github.com/pkg/errors"

type Driver interface {
	Open() error
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error

	BatchGet(keys [][]byte) ([][]byte, error)
	BatchPut(keys, values [][]byte) error
	BatchDelete(keys [][]byte) error

	Has(key []byte) bool
	Close() error

	Path() string
	Fold(func(k, v []byte) error) int64
	Keys(func(k []byte) error) int64
}

var (
	ErrNumKeyValueNotMatch = errors.New("number of key and value does not match")
)

func New(path string) (Driver, error) {
	return newBadgerDriver(path)
}
