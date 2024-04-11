package bus

import (
	"context"
	"gar/app/ent"
	"gar/app/search"
	"gar/searcher"
)

type Context struct {
	Upstream context.Context
	Indexer  *searcher.Indexer
	Request  *search.RequestBody
	Results  []*ent.BiliBiliVideo
}

type UN string
