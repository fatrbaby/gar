package engine

import (
	"gar/ent"
)

type Index interface {
	Add(doc *ent.Document)
	Delete(uid uint64, keyword *ent.Keyword)
	Search(q *ent.TermQuery, onFlag uint64, offFlag uint64, orFlags []uint64) []string
}

type Value struct {
	Id      string
	Feature uint64
}
