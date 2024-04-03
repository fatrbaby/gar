package db

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"os"
	"path"
	"sync/atomic"
)

type BadgerDriver struct {
	db   *badger.DB
	path string
}

func newBadgerDriver(path string) (Driver, error) {
	b := &BadgerDriver{}
	b.WithPath(path)
	if err := b.Open(); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *BadgerDriver) WithPath(path string) *BadgerDriver {
	b.path = path
	return b
}

func (b *BadgerDriver) Open() error {
	dir := b.Path()

	if err := os.MkdirAll(path.Base(dir), os.ModePerm); err != nil {
		return err
	}

	options := badger.DefaultOptions(dir).WithNumVersionsToKeep(1).WithLoggingLevel(badger.ERROR)
	db, err := badger.Open(options)

	if err != nil {
		return err
	}
	b.db = db

	return nil
}

func (b *BadgerDriver) Put(key, value []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

func (b *BadgerDriver) Get(key []byte) ([]byte, error) {
	var value []byte

	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			value = val
			return nil
		})
	})

	return value, err
}

func (b *BadgerDriver) Delete(key []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func (b *BadgerDriver) BatchGet(keys [][]byte) ([][]byte, error) {
	txn := b.db.NewTransaction(false)
	values := make([][]byte, len(keys))

	var err error
	var item *badger.Item

	for i, key := range keys {
		item, err = txn.Get(key)
		if err == nil {
			var value []byte
			err = item.Value(func(val []byte) error {
				value = val
				return nil
			})

			if err == nil {
				values[i] = value
			} else {
				values[i] = []byte{}
			}
		} else {
			values[i] = []byte{}
			if !errors.Is(err, badger.ErrKeyNotFound) {
				txn.Discard()
				txn = b.db.NewTransaction(false)
			}
		}
	}

	txn.Discard()

	return values, err
}

func (b *BadgerDriver) BatchPut(keys, values [][]byte) error {
	if len(keys) != len(values) {
		return ErrNumKeyValueNotMatch
	}

	var err error

	txn := b.db.NewTransaction(true)

	for i, key := range keys {
		value := values[i]
		if err = txn.Set(key, value); err != nil {
			_ = txn.Commit()
			txn = b.db.NewTransaction(true)
			_ = txn.Set(key, value)
		}
	}

	_ = txn.Commit()

	return err
}

func (b *BadgerDriver) BatchDelete(keys [][]byte) error {
	var err error
	txn := b.db.NewTransaction(true)

	for _, key := range keys {
		if err = txn.Delete(key); err != nil {
			_ = txn.Commit()
			txn = b.db.NewTransaction(true)
			_ = txn.Delete(key)
		}
	}

	_ = txn.Commit()

	return err
}

func (b *BadgerDriver) Has(key []byte) bool {
	var has bool

	_ = b.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get(key)
		if err != nil {
			return err
		}

		has = true
		return nil
	})

	return has
}

func (b *BadgerDriver) Close() error {
	return b.db.Close()
}

func (b *BadgerDriver) Path() string {
	return b.path
}

func (b *BadgerDriver) Fold(f func(k []byte, v []byte) error) int64 {
	var num int64

	_ = b.db.View(func(txn *badger.Txn) error {
		iter := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iter.Close()

		for iter.Rewind(); iter.Valid(); iter.Next() {
			item := iter.Item()
			var value []byte

			err := item.Value(func(val []byte) error {
				value = val
				return nil
			})

			if err != nil {
				continue
			}

			if err := f(item.Key(), value); err != nil {
				atomic.AddInt64(&num, 1)
			}

			return nil
		}

		return nil
	})

	return atomic.LoadInt64(&num)
}

func (b *BadgerDriver) Keys(f func(k []byte) error) int64 {
	var num int64

	_ = b.db.View(func(txn *badger.Txn) error {
		options := badger.DefaultIteratorOptions
		// fetch key only
		options.PrefetchValues = false

		iter := txn.NewIterator(options)
		defer iter.Close()

		for iter.Rewind(); iter.Valid(); iter.Next() {
			if err := f(iter.Item().Key()); err == nil {
				atomic.AddInt64(&num, 1)
			}
		}

		return nil
	})

	return atomic.LoadInt64(&num)
}

func (b *BadgerDriver) Gc() {
	lsmSize1, vlogSize1 := b.db.Size()

	for {
		err := b.db.RunValueLogGC(0.5)

		if errors.Is(err, badger.ErrNoRewrite) || errors.Is(err, badger.ErrRejected) {
			break
		}
	}

	lsmSize2, vlogSize2 := b.db.Size()

	if vlogSize2 < vlogSize1 {
		slog.Info(
			"badger before GC, LSM {}, vlog {}. after GC, LSM {}, vlog {}",
			lsmSize1,
			vlogSize1,
			lsmSize2,
			vlogSize2,
		)
	} else {
		slog.Info("collect zero garbage")
	}
}
