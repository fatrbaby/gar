package bus

import (
	"context"
	"gar/app/ent"
	"gar/searcher"
)

type Context struct {
	Upstream context.Context
	Indexer  *searcher.Indexer
	Request  *ent.RequestBody
	Results  []*ent.BiliBiliVideo
}

type UN string
