package storage

import (
	"gar/app/data"
	"gar/searcher"
)

type Worker interface {
	WithDatasource(data.Source) Worker
	Run()
	Indexer() *searcher.Indexer
}
