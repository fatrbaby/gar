package shortcut

import (
	hasher "github.com/leemcloughlin/gofarmhash"
	"golang.org/x/exp/maps"
	"sync"
)

type ConcurrentHashMap struct {
	buckets  []map[string]any
	segments int
	locks    []sync.RWMutex
	seeds    uint32
}

type MapEntry struct {
	Key   string
	Value any
}

type MapIterator interface {
	Next() *MapEntry
}

type ConcurrentHashMapIterator struct {
	c           *ConcurrentHashMap
	keys        [][]string
	rowIndex    int
	columnIndex int
}

// NewConcurrentHashMap 创建一个并发安全的哈希表
func NewConcurrentHashMap(segments, cap int) *ConcurrentHashMap {
	buckets := make([]map[string]any, segments)
	locks := make([]sync.RWMutex, segments)

	for i := 0; i < segments; i++ {
		buckets[i] = make(map[string]any, cap/segments)
	}

	return &ConcurrentHashMap{
		buckets:  buckets,
		segments: segments,
		locks:    locks,
		seeds:    0,
	}
}

// Set 设置键值对
func (c *ConcurrentHashMap) Set(key string, value any) {
	idx := c.getSegmentIndex(key)
	c.locks[idx].RLock()
	defer c.locks[idx].RUnlock()

	c.buckets[idx][key] = value
}

// Get 获取键值对
func (c *ConcurrentHashMap) Get(key string) (any, bool) {
	idx := c.getSegmentIndex(key)
	c.locks[idx].RLock()
	defer c.locks[idx].RUnlock()

	value, has := c.buckets[idx][key]

	return value, has
}

func (c *ConcurrentHashMap) Iterator() *ConcurrentHashMapIterator {
	keys := make([][]string, len(c.buckets))

	for _, m := range c.buckets {
		row := maps.Keys(m)
		keys = append(keys, row)
	}

	return &ConcurrentHashMapIterator{
		c:    c,
		keys: keys,
	}
}

// 获取分段索引
func (c *ConcurrentHashMap) getSegmentIndex(key string) int {
	seed := hasher.Hash32WithSeed([]byte(key), c.seeds)

	return int(seed) % c.segments
}

func (i *ConcurrentHashMapIterator) Next() *MapEntry {
	if i.rowIndex >= len(i.keys) {
		return nil
	}

	row := i.keys[i.columnIndex]

	if len(row) == 0 {
		i.rowIndex++
		return i.Next()
	}

	key := row[i.columnIndex]
	value, _ := i.c.Get(key)

	if i.columnIndex >= len(row)-1 {
		i.rowIndex++
		i.columnIndex = 0
	} else {
		i.columnIndex++
	}

	return &MapEntry{
		Key:   key,
		Value: value,
	}
}
