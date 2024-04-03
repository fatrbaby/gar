package engine

import (
	"gar/ent"
	"gar/shortcut"
	"github.com/huandu/skiplist"
	hasher "github.com/leemcloughlin/gofarmhash"
	"runtime"
	"sync"
)

const capLock = 1000

// ReverseIndex 倒排索引
type ReverseIndex struct {
	table *shortcut.ConcurrentHashMap
	locks []sync.RWMutex
}

func NewReverseIndexer(docEstimate int) *ReverseIndex {
	return &ReverseIndex{
		table: shortcut.NewConcurrentHashMap(runtime.NumCPU(), docEstimate),
		locks: make([]sync.RWMutex, capLock),
	}
}

func (i *ReverseIndex) Add(doc *ent.Document) {
	for _, keyword := range doc.Keywords {
		key := keyword.IntoString()
		lock := i.getLock(key)
		lock.Lock()
		newValue := Value{Id: doc.Id, Feature: doc.Features}

		if value, has := i.table.Get(key); has {
			list := value.(*skiplist.SkipList)
			list.Set(doc.Uid, newValue)
		} else {
			list := skiplist.New(skiplist.Uint64)
			list.Set(doc.Uid, newValue)
			i.table.Set(key, list)
		}
		lock.Unlock()
	}
}

func (i *ReverseIndex) Delete(uid uint64, keyword *ent.Keyword) {
	key := keyword.IntoString()
	lock := i.getLock(key)
	lock.Lock()

	if value, has := i.table.Get(key); has {
		list := value.(*skiplist.SkipList)
		list.Remove(uid)
	}

	lock.Unlock()
}

func (i *ReverseIndex) Search(q *ent.TermQuery, onFlag uint64, offFlag uint64, orFlags []uint64) []string {
	result := i.search(q, onFlag, offFlag, orFlags)

	if result == nil {
		return nil
	}

	container := make([]string, 0, result.Len())
	node := result.Front()

	for node != nil {
		value := node.Value.(Value)
		container = append(container, value.Id)
		node = node.Next()
	}

	return container
}

func (i *ReverseIndex) search(q *ent.TermQuery, onFlag uint64, offFlag uint64, orFlags []uint64) *skiplist.SkipList {
	if q.Keyword != nil {
		kw := q.Keyword.IntoString()
		if value, has := i.table.Get(kw); has {
			r := skiplist.New(skiplist.Uint64)
			list := value.(*skiplist.SkipList)
			node := list.Front()

			if node != nil {
				uid := node.Key().(uint64)
				newValue, _ := node.Value.(Value)
				flag := newValue.Feature
				if uid > 0 && i.filterByBits(flag, onFlag, offFlag, orFlags) {
					r.Set(uid, newValue)
				}
				node = node.Next()
			}

			return r
		}
	} else if len(q.Must) > 0 {
		r := make([]*skiplist.SkipList, 0, len(q.Must))

		for _, must := range q.Must {
			r = append(r, i.search(must, onFlag, offFlag, orFlags))
		}

		return shortcut.Intersection(r...)
	} else if len(q.Should) > 0 {
		r := make([]*skiplist.SkipList, 0, len(q.Should))
		for _, should := range q.Should {
			r = append(r, i.search(should, onFlag, offFlag, orFlags))
		}

		return shortcut.UnionSet(r...)
	}

	return nil
}

func (i *ReverseIndex) filterByBits(bits uint64, onFlag uint64, offFlag uint64, orFlags []uint64) bool {
	if bits&onFlag != onFlag {
		return false
	}

	return true
}

func (i *ReverseIndex) getLock(key string) sync.RWMutex {
	seed := hasher.Hash32WithSeed([]byte(key), 0)

	return i.locks[seed%capLock]
}
