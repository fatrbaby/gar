package searcher

import (
	"bytes"
	"encoding/gob"
	"gar/db"
	"gar/engine"
	"gar/ent"
	"log/slog"
	"strings"
	"sync/atomic"
)

type Indexer struct {
	forwardIndex db.Driver
	reverseIndex engine.Index
	maxUid       uint64
}

func NewIndexer() *Indexer {
	return &Indexer{}
}

func (i *Indexer) Setup(docEstimate int, path string) error {
	driver, err := db.New(path)

	if err != nil {
		return err
	}

	i.forwardIndex = driver
	i.reverseIndex = engine.NewReverseIndexer(docEstimate)

	return nil
}

func (i *Indexer) Search(q *ent.TermQuery, onFlag uint64, offFlag uint64, orFlags []uint64) []*ent.Document {
	docIds := i.reverseIndex.Search(q, onFlag, offFlag, orFlags)

	if len(docIds) == 0 {
		return nil
	}

	keys := make([][]byte, 0, len(docIds))

	for _, docId := range docIds {
		keys = append(keys, []byte(docId))
	}

	docs, err := i.forwardIndex.BatchGet(keys)

	if err != nil {
		slog.Error("db read failed", "error", err)
		return nil
	}

	documents := make([]*ent.Document, 0, len(docs))

	for _, doc := range docs {
		if len(doc) > 0 {
			reader := bytes.NewReader([]byte{})
			reader.Reset(doc)
			decoder := gob.NewDecoder(reader)
			var d ent.Document

			if err := decoder.Decode(&d); err == nil {
				documents = append(documents, &d)
			}
		}
	}

	return documents
}

func (i *Indexer) LoadFromFile() int {
	reader := bytes.NewReader([]byte{})

	n := i.forwardIndex.Fold(func(k, v []byte) error {
		reader.Reset(v)
		decoder := gob.NewDecoder(reader)
		var doc ent.Document

		if err := decoder.Decode(&doc); err != nil {
			slog.Warn("gob decode failed", "error", err)
			return err
		}

		i.reverseIndex.Add(&doc)

		return nil
	})

	slog.Info("loaded data from forward index", "count", n, "path", i.forwardIndex.Path())

	return int(n)
}

func (i *Indexer) Add(doc *ent.Document) (int, error) {
	docId := strings.TrimSpace(doc.Id)

	if len(docId) == 0 {
		return 0, nil
	}

	i.Delete(docId)

	doc.Uid = atomic.AddUint64(&i.maxUid, 1)

	// write forward index
	var value bytes.Buffer
	encoder := gob.NewEncoder(&value)

	if err := encoder.Encode(doc); err != nil {
		return 0, err
	}

	err := i.forwardIndex.Put([]byte(docId), value.Bytes())

	if err != nil {
		return 0, err
	}

	i.reverseIndex.Add(doc)

	return 1, nil
}

func (i *Indexer) Delete(id string) int {
	n := 0
	forwardKey := []byte(id)
	content, err := i.forwardIndex.Get(forwardKey)

	if err == nil && len(content) > 0 {
		n = 1
		reader := bytes.NewReader([]byte{})
		reader.Reset(content)
		decoder := gob.NewDecoder(reader)
		var doc ent.Document

		if err := decoder.Decode(&doc); err == nil {
			for _, keyword := range doc.Keywords {
				i.reverseIndex.Delete(doc.Uid, keyword)
			}
		}
	}

	return n
}

func (i *Indexer) Count() int {
	n := 0

	i.forwardIndex.Keys(func(k []byte) error {
		n++
		return nil
	})

	return n
}

func (i *Indexer) Close() error {
	return i.forwardIndex.Close()
}
